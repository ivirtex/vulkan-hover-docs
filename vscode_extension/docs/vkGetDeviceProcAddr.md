# vkGetDeviceProcAddr(3) Manual Page

## Name

vkGetDeviceProcAddr - Return a function pointer for a command



## [](#_c_specification)C Specification

In order to support systems with multiple Vulkan implementations, the function pointers returned by [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html) **may** point to dispatch code that calls a different real implementation for different [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) objects or their child objects. The overhead of the internal dispatch for [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) objects can be avoided by obtaining device-specific function pointers for any commands that use a device or device-child object as their dispatchable object. Such function pointers **can** be obtained by calling:

```c++
// Provided by VK_VERSION_1_0
PFN_vkVoidFunction vkGetDeviceProcAddr(
    VkDevice                                    device,
    const char*                                 pName);
```

## [](#_description)Description

The table below defines the various use cases for `vkGetDeviceProcAddr` and expected return value (“fp” is “function pointer”) for each case. A valid returned function pointer (“fp”) **must** not be `NULL`.

The returned function pointer is of type [PFN\_vkVoidFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkVoidFunction.html), and **must** be cast to the type of the command being queried before use. The function pointer **must** only be called with a dispatchable object (the first parameter) that is `device` or a child of `device`.

Table 1. `vkGetDeviceProcAddr` behavior    `device` `pName` return value

`NULL`

\*1

undefined

invalid device

\*1

undefined

device

`NULL`

undefined

device

requested core version2 device-level dispatchable command3

fp4

device

enabled extension device-level dispatchable command3

fp4

any other case, not covered above

`NULL`

1

"\*" means any representable value for the parameter (including valid values, invalid values, and `NULL`).

2

Device-level commands which are part of the core version specified by [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` when creating the instance will always return a valid function pointer. If the [`maintenance5`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance5) feature is enabled, core commands beyond that version which are supported by the implementation will return `NULL`, otherwise the implementation **may** either return `NULL` or a function pointer. If a function pointer is returned, it **must** not be called.

3

In this function, device-level excludes all physical-device-level commands.

4

The returned function pointer **must** only be called with a dispatchable object (the first parameter) that is `device` or a child of `device` e.g. [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html), or [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html).

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceProcAddr-device-parameter)VUID-vkGetDeviceProcAddr-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceProcAddr-pName-parameter)VUID-vkGetDeviceProcAddr-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[PFN\_vkVoidFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkVoidFunction.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceProcAddr)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0