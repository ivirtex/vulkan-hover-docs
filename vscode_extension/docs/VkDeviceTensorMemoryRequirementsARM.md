# VkDeviceTensorMemoryRequirementsARM(3) Manual Page

## Name

VkDeviceTensorMemoryRequirementsARM - (None)



## [](#_c_specification)C Specification

The `VkDeviceTensorMemoryRequirementsARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkDeviceTensorMemoryRequirementsARM {
    VkStructureType                 sType;
    const void*                     pNext;
    const VkTensorCreateInfoARM*    pCreateInfo;
} VkDeviceTensorMemoryRequirementsARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pCreateInfo` is a pointer to a [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html) structure containing parameters affecting the creation of the tensor to query.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDeviceTensorMemoryRequirementsARM-sType-sType)VUID-VkDeviceTensorMemoryRequirementsARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_TENSOR_MEMORY_REQUIREMENTS_ARM`
- [](#VUID-VkDeviceTensorMemoryRequirementsARM-pNext-pNext)VUID-VkDeviceTensorMemoryRequirementsARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDeviceTensorMemoryRequirementsARM-pCreateInfo-parameter)VUID-VkDeviceTensorMemoryRequirementsARM-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html) structure

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html), [vkGetDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceTensorMemoryRequirementsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceTensorMemoryRequirementsARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0