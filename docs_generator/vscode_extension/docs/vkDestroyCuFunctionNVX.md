# vkDestroyCuFunctionNVX(3) Manual Page

## Name

vkDestroyCuFunctionNVX - Stub description of vkDestroyCuFunctionNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
void vkDestroyCuFunctionNVX(
    VkDevice                                    device,
    VkCuFunctionNVX                             function,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDestroyCuFunctionNVX-device-parameter)VUID-vkDestroyCuFunctionNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyCuFunctionNVX-function-parameter)VUID-vkDestroyCuFunctionNVX-function-parameter  
  `function` **must** be a valid [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html) handle
- [](#VUID-vkDestroyCuFunctionNVX-pAllocator-parameter)VUID-vkDestroyCuFunctionNVX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyCuFunctionNVX-function-parent)VUID-vkDestroyCuFunctionNVX-function-parent  
  `function` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyCuFunctionNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0