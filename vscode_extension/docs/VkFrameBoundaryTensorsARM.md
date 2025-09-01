# VkFrameBoundaryTensorsARM(3) Manual Page

## Name

VkFrameBoundaryTensorsARM - Add tensor frame boundary information to queue submissions



## [](#_c_specification)C Specification

The `VkFrameBoundaryTensorsARM` structure is defined as:

```c++
// Provided by VK_EXT_frame_boundary with VK_ARM_tensors
typedef struct VkFrameBoundaryTensorsARM {
    VkStructureType       sType;
    const void*           pNext;
    uint32_t              tensorCount;
    const VkTensorARM*    pTensors;
} VkFrameBoundaryTensorsARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensorCount` is the number of tensors that store frame results.
- `pTensors` is a pointer to an array of [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) objects with tensorCount entries.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkFrameBoundaryTensorsARM-sType-sType)VUID-VkFrameBoundaryTensorsARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAME_BOUNDARY_TENSORS_ARM`
- [](#VUID-VkFrameBoundaryTensorsARM-pTensors-parameter)VUID-VkFrameBoundaryTensorsARM-pTensors-parameter  
  `pTensors` **must** be a valid pointer to an array of `tensorCount` valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handles
- [](#VUID-VkFrameBoundaryTensorsARM-tensorCount-arraylength)VUID-VkFrameBoundaryTensorsARM-tensorCount-arraylength  
  `tensorCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_EXT\_frame\_boundary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_frame_boundary.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFrameBoundaryTensorsARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0