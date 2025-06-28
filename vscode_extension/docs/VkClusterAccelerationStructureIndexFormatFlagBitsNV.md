# VkClusterAccelerationStructureIndexFormatFlagBitsNV(3) Manual Page

## Name

VkClusterAccelerationStructureIndexFormatFlagBitsNV - Bits specifying the index type in the index buffer



## [](#_c_specification)C Specification

Bits that **can** be set in [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)::`indexType`, [VkClusterAccelerationStructureBuildTriangleClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterInfoNV.html)::`opacityMicromapIndexType`, [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)::`indexType` and [VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureBuildTriangleClusterTemplateInfoNV.html)::`opacityMicromapIndexType` specifying the index type is one of:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef enum VkClusterAccelerationStructureIndexFormatFlagBitsNV {
    VK_CLUSTER_ACCELERATION_STRUCTURE_INDEX_FORMAT_8BIT_NV = 0x00000001,
    VK_CLUSTER_ACCELERATION_STRUCTURE_INDEX_FORMAT_16BIT_NV = 0x00000002,
    VK_CLUSTER_ACCELERATION_STRUCTURE_INDEX_FORMAT_32BIT_NV = 0x00000004,
} VkClusterAccelerationStructureIndexFormatFlagBitsNV;
```

## [](#_description)Description

- `VK_CLUSTER_ACCELERATION_STRUCTURE_INDEX_FORMAT_8BIT_NV` specifies that 8-bit indices will be used.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_INDEX_FORMAT_16BIT_NV` specifies that 16-bit indices will be used.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_INDEX_FORMAT_32BIT_NV` specifies that 32-bit indices will be used.

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureIndexFormatFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureIndexFormatFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureIndexFormatFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0