# VkMemoryDedicatedAllocateInfoTensorARM(3) Manual Page

## Name

VkMemoryDedicatedAllocateInfoTensorARM - Specify a dedicated memory allocation tensor resource



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkMemoryDedicatedAllocateInfoTensorARM` structure, then that structure includes a handle of the sole tensor resource that the memory **can** be bound to.

The `VkMemoryDedicatedAllocateInfoTensorARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkMemoryDedicatedAllocateInfoTensorARM {
    VkStructureType    sType;
    const void*        pNext;
    VkTensorARM        tensor;
} VkMemoryDedicatedAllocateInfoTensorARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensor` is a handle of a tensor which this memory will be bound to.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryDedicatedAllocateInfoTensorARM-allocationSize-09710)VUID-VkMemoryDedicatedAllocateInfoTensorARM-allocationSize-09710  
  `VkMemoryAllocateInfo`::`allocationSize` **must** equal the `VkMemoryRequirements`::`size` of the tensor
- [](#VUID-VkMemoryDedicatedAllocateInfoTensorARM-tensor-09859)VUID-VkMemoryDedicatedAllocateInfoTensorARM-tensor-09859  
  If [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`, the memory being imported **must** also be a dedicated tensor allocation and `tensor` **must** be identical to the tensor associated with the imported memory

Valid Usage (Implicit)

- [](#VUID-VkMemoryDedicatedAllocateInfoTensorARM-sType-sType)VUID-VkMemoryDedicatedAllocateInfoTensorARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_TENSOR_ARM`
- [](#VUID-VkMemoryDedicatedAllocateInfoTensorARM-tensor-parameter)VUID-VkMemoryDedicatedAllocateInfoTensorARM-tensor-parameter  
  `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryDedicatedAllocateInfoTensorARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0