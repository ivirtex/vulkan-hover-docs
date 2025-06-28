# vkDestroyCuModuleNVX(3) Manual Page

## Name

vkDestroyCuModuleNVX - Stub description of vkDestroyCuModuleNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
void vkDestroyCuModuleNVX(
    VkDevice                                    device,
    VkCuModuleNVX                               module,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDestroyCuModuleNVX-device-parameter)VUID-vkDestroyCuModuleNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyCuModuleNVX-module-parameter)VUID-vkDestroyCuModuleNVX-module-parameter  
  `module` **must** be a valid [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html) handle
- [](#VUID-vkDestroyCuModuleNVX-pAllocator-parameter)VUID-vkDestroyCuModuleNVX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyCuModuleNVX-module-parent)VUID-vkDestroyCuModuleNVX-module-parent  
  `module` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyCuModuleNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0