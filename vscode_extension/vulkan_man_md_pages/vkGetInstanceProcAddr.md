# vkGetInstanceProcAddr(3) Manual Page

## Name

vkGetInstanceProcAddr - Return a function pointer for a command



## <a href="#_c_specification" class="anchor"></a>C Specification

Function pointers for all Vulkan commands **can** be obtained by
calling:

``` c
// Provided by VK_VERSION_1_0
PFN_vkVoidFunction vkGetInstanceProcAddr(
    VkInstance                                  instance,
    const char*                                 pName);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance that the function pointer will be
  compatible with, or `NULL` for commands not dependent on any instance.

- `pName` is the name of the command to obtain.

## <a href="#_description" class="anchor"></a>Description

`vkGetInstanceProcAddr` itself is obtained in a platform- and loader-
specific manner. Typically, the loader library will export this command
as a function symbol, so applications **can** link against the loader
library, or load it dynamically and look up the symbol using
platform-specific APIs.

The table below defines the various use cases for
`vkGetInstanceProcAddr` and expected return value (“fp” is “function
pointer”) for each case. A valid returned function pointer (“fp”)
**must** not be `NULL`.

The returned function pointer is of type
[PFN_vkVoidFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkVoidFunction.html), and **must** be cast to
the type of the command being queried before use.

| `instance` | `pName` | return value |
|----|----|----|
| \*<sup>1</sup> | `NULL` | undefined |
| invalid non-`NULL` instance | \*<sup>1</sup> | undefined |
| `NULL` | *global command*<sup>2</sup> | fp |
| `NULL` | [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetInstanceProcAddr.html) | fp<sup>5</sup> |
| instance | [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetInstanceProcAddr.html) | fp |
| instance | core *dispatchable command* | fp<sup>3</sup> |
| instance | enabled instance extension dispatchable command for `instance` | fp<sup>3</sup> |
| instance | available device extension<sup>4</sup> dispatchable command for `instance` | fp<sup>3</sup> |
| any other case, not covered above |  | `NULL` |

Table 1. `vkGetInstanceProcAddr` behavior

1  
"\*" means any representable value for the parameter (including valid
values, invalid values, and `NULL`).

2  
The global commands are:
[vkEnumerateInstanceVersion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceVersion.html),
[vkEnumerateInstanceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceExtensionProperties.html),
[vkEnumerateInstanceLayerProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateInstanceLayerProperties.html),
and [vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html). Dispatchable commands are
all other commands which are not global.

3  
The returned function pointer **must** only be called with a
dispatchable object (the first parameter) that is `instance` or a child
of `instance`, e.g. [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html), or [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html).

4  
An “available device extension” is a device extension supported by any
physical device enumerated by `instance`.

5  
Starting with Vulkan 1.2, `vkGetInstanceProcAddr` can resolve itself
with a `NULL` instance pointer.

Valid Usage (Implicit)

- <a href="#VUID-vkGetInstanceProcAddr-instance-parameter"
  id="VUID-vkGetInstanceProcAddr-instance-parameter"></a>
  VUID-vkGetInstanceProcAddr-instance-parameter  
  If `instance` is not `NULL`, `instance` **must** be a valid
  [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkGetInstanceProcAddr-pName-parameter"
  id="VUID-vkGetInstanceProcAddr-pName-parameter"></a>
  VUID-vkGetInstanceProcAddr-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkVoidFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkVoidFunction.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetInstanceProcAddr"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
