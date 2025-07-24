# vkCreateCuFunctionNVX(3) Manual Page

## Name

vkCreateCuFunctionNVX - Stub description of vkCreateCuFunctionNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_NVX_binary_import
VkResult vkCreateCuFunctionNVX(
    VkDevice                                    device,
    const VkCuFunctionCreateInfoNVX*            pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkCuFunctionNVX*                            pFunction);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateCuFunctionNVX-device-parameter)VUID-vkCreateCuFunctionNVX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateCuFunctionNVX-pCreateInfo-parameter)VUID-vkCreateCuFunctionNVX-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkCuFunctionCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionCreateInfoNVX.html) structure
- [](#VUID-vkCreateCuFunctionNVX-pAllocator-parameter)VUID-vkCreateCuFunctionNVX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateCuFunctionNVX-pFunction-parameter)VUID-vkCreateCuFunctionNVX-pFunction-parameter  
  `pFunction` **must** be a valid pointer to a [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html) handle
- [](#VUID-vkCreateCuFunctionNVX-device-queuecount)VUID-vkCreateCuFunctionNVX-device-queuecount  
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

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkCuFunctionCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionCreateInfoNVX.html), [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateCuFunctionNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0