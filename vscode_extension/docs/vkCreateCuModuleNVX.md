# vkCreateCuModuleNVX(3) Manual Page

## Name

vkCreateCuModuleNVX - Stub description of vkCreateCuModuleNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
VkResult vkCreateCuModuleNVX(
    VkDevice                                    device,
    const VkCuModuleCreateInfoNVX*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCuModuleNVX*                              pModule);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateCuModuleNVX-device-parameter)VUID-vkCreateCuModuleNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateCuModuleNVX-pCreateInfo-parameter)VUID-vkCreateCuModuleNVX-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleCreateInfoNVX.html) structure
- [](#VUID-vkCreateCuModuleNVX-pAllocator-parameter)VUID-vkCreateCuModuleNVX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateCuModuleNVX-pModule-parameter)VUID-vkCreateCuModuleNVX-pModule-parameter  
  `pModule` **must** be a valid pointer to a [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html) handle
- [](#VUID-vkCreateCuModuleNVX-device-queuecount)VUID-vkCreateCuModuleNVX-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleCreateInfoNVX.html), [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateCuModuleNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0