# VkTensorMemoryRequirementsInfoARM(3) Manual Page

## Name

VkTensorMemoryRequirementsInfoARM - Structure specifying memory requirements



## [](#_c_specification)C Specification

The `VkTensorMemoryRequirementsInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorMemoryRequirementsInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkTensorARM        tensor;
} VkTensorMemoryRequirementsInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensor` is the tensor to query.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkTensorMemoryRequirementsInfoARM-sType-sType)VUID-VkTensorMemoryRequirementsInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_MEMORY_REQUIREMENTS_INFO_ARM`
- [](#VUID-VkTensorMemoryRequirementsInfoARM-pNext-pNext)VUID-VkTensorMemoryRequirementsInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkTensorMemoryRequirementsInfoARM-tensor-parameter)VUID-VkTensorMemoryRequirementsInfoARM-tensor-parameter  
  `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [vkGetTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorMemoryRequirementsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorMemoryRequirementsInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0