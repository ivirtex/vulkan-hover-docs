# vkGetTensorMemoryRequirementsARM(3) Manual Page

## Name

vkGetTensorMemoryRequirementsARM - Returns the memory requirements for specified Vulkan object



## [](#_c_specification)C Specification

To determine the memory requirements for a tensor resource, call:

```c++
// Provided by VK_ARM_tensors
void vkGetTensorMemoryRequirementsARM(
    VkDevice                                    device,
    const VkTensorMemoryRequirementsInfoARM*    pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the tensor.
- `pInfo` is a pointer to a [VkTensorMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryRequirementsInfoARM.html) structure containing parameters required for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the tensor object are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetTensorMemoryRequirementsARM-device-parameter)VUID-vkGetTensorMemoryRequirementsARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetTensorMemoryRequirementsARM-pInfo-parameter)VUID-vkGetTensorMemoryRequirementsARM-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkTensorMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryRequirementsInfoARM.html) structure
- [](#VUID-vkGetTensorMemoryRequirementsARM-pMemoryRequirements-parameter)VUID-vkGetTensorMemoryRequirementsARM-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html), [VkTensorMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryRequirementsInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetTensorMemoryRequirementsARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0