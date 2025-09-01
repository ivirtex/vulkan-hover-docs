# VkTensorDependencyInfoARM(3) Manual Page

## Name

VkTensorDependencyInfoARM - Structure specifying tensor dependency information for a synchronization command



## [](#_c_specification)C Specification

The `VkTensorDependencyInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorDependencyInfoARM {
    VkStructureType                    sType;
    const void*                        pNext;
    uint32_t                           tensorMemoryBarrierCount;
    const VkTensorMemoryBarrierARM*    pTensorMemoryBarriers;
} VkTensorDependencyInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensorMemoryBarrierCount` is the length of the `pTensorMemoryBarriers` array.
- `pTensorMemoryBarriers` is a pointer to an array of [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html) structures defining memory dependencies between tensors.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkTensorDependencyInfoARM-sType-sType)VUID-VkTensorDependencyInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_DEPENDENCY_INFO_ARM`
- [](#VUID-VkTensorDependencyInfoARM-pTensorMemoryBarriers-parameter)VUID-VkTensorDependencyInfoARM-pTensorMemoryBarriers-parameter  
  `pTensorMemoryBarriers` **must** be a valid pointer to a valid [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html) structure

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorDependencyInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0