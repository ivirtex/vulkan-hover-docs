# vkGetDeviceImageSparseMemoryRequirements(3) Manual Page

## Name

vkGetDeviceImageSparseMemoryRequirements - Query the memory requirements for a sparse image



## [](#_c_specification)C Specification

To determine the sparse memory requirements for an image resource without creating an object, call:

```c++
// Provided by VK_VERSION_1_3
void vkGetDeviceImageSparseMemoryRequirements(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance4
void vkGetDeviceImageSparseMemoryRequirementsKHR(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device intended to own the image.
- `pInfo` is a pointer to a [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html) structure containing parameters required for the memory requirements query.
- `pSparseMemoryRequirementCount` is a pointer to an integer related to the number of sparse memory requirements available or queried, as described below.
- `pSparseMemoryRequirements` is either `NULL` or a pointer to an array of `VkSparseImageMemoryRequirements2` structures.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceImageSparseMemoryRequirements-device-parameter)VUID-vkGetDeviceImageSparseMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceImageSparseMemoryRequirements-pInfo-parameter)VUID-vkGetDeviceImageSparseMemoryRequirements-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html) structure
- [](#VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter)VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter  
  `pSparseMemoryRequirementCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter)VUID-vkGetDeviceImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter  
  If the value referenced by `pSparseMemoryRequirementCount` is not `0`, and `pSparseMemoryRequirements` is not `NULL`, `pSparseMemoryRequirements` **must** be a valid pointer to an array of `pSparseMemoryRequirementCount` [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements2.html) structures

## [](#_see_also)See Also

[VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html), [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceImageSparseMemoryRequirements).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0