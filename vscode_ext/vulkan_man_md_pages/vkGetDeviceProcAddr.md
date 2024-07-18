# vkGetDeviceProcAddr(3) Manual Page

## Name

vkGetDeviceProcAddr - Return a function pointer for a command



## <a href="#_c_specification" class="anchor"></a>C Specification

In order to support systems with multiple Vulkan implementations, the
function pointers returned by
[vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetInstanceProcAddr.html) **may** point to
dispatch code that calls a different real implementation for different
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) objects or their child objects. The overhead
of the internal dispatch for [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) objects can be
avoided by obtaining device-specific function pointers for any commands
that use a device or device-child object as their dispatchable object.
Such function pointers **can** be obtained by calling:

``` c
// Provided by VK_VERSION_1_0
PFN_vkVoidFunction vkGetDeviceProcAddr(
    VkDevice                                    device,
    const char*                                 pName);
```

## <a href="#_description" class="anchor"></a>Description

The table below defines the various use cases for `vkGetDeviceProcAddr`
and expected return value (“fp” is “function pointer”) for each case. A
valid returned function pointer (“fp”) **must** not be `NULL`.

The returned function pointer is of type
[PFN_vkVoidFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkVoidFunction.html), and **must** be cast to
the type of the command being queried before use. The function pointer
**must** only be called with a dispatchable object (the first parameter)
that is `device` or a child of `device`.

| `device` | `pName` | return value |
|----|----|----|
| `NULL` | \*<sup>1</sup> | undefined |
| invalid device | \*<sup>1</sup> | undefined |
| device | `NULL` | undefined |
| device | requested core version<sup>2</sup> device-level dispatchable command<sup>3</sup> | fp<sup>4</sup> |
| device | enabled extension device-level dispatchable command<sup>3</sup> | fp<sup>4</sup> |
| any other case, not covered above |  | `NULL` |

Table 1. `vkGetDeviceProcAddr` behavior

1  
"\*" means any representable value for the parameter (including valid
values, invalid values, and `NULL`).

2  
Device-level commands which are part of the core version specified by
[VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` when creating
the instance will always return a valid function pointer. If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
target="_blank" rel="noopener"><code>maintenance5</code></a> feature is
enabled, core commands beyond that version which are supported by the
implementation will return `NULL`, otherwise the implementation **may**
either return `NULL` or a function pointer. If a function pointer is
returned, it **must** not be called.

3  
In this function, device-level excludes all physical-device-level
commands.

4  
The returned function pointer **must** only be called with a
dispatchable object (the first parameter) that is `device` or a child of
`device` e.g. [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html), or
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html).

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceProcAddr-device-parameter"
  id="VUID-vkGetDeviceProcAddr-device-parameter"></a>
  VUID-vkGetDeviceProcAddr-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceProcAddr-pName-parameter"
  id="VUID-vkGetDeviceProcAddr-pName-parameter"></a>
  VUID-vkGetDeviceProcAddr-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkVoidFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkVoidFunction.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceProcAddr"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
