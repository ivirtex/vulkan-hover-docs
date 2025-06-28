# VkClusterAccelerationStructureAddressResolutionFlagBitsNV(3) Manual Page

## Name

VkClusterAccelerationStructureAddressResolutionFlagBitsNV - Bitmask specifying address resolution flags in cluster acceleration structure



## [](#_c_specification)C Specification

Bits which **can** be set in [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`addressResolutionFlags`, specifying how the device address in [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html) are interpreted, are:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef enum VkClusterAccelerationStructureAddressResolutionFlagBitsNV {
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_NONE_NV = 0,
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_DST_IMPLICIT_DATA_BIT_NV = 0x00000001,
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_SCRATCH_DATA_BIT_NV = 0x00000002,
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_DST_ADDRESS_ARRAY_BIT_NV = 0x00000004,
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_DST_SIZES_ARRAY_BIT_NV = 0x00000008,
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_SRC_INFOS_ARRAY_BIT_NV = 0x00000010,
    VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_SRC_INFOS_COUNT_BIT_NV = 0x00000020,
} VkClusterAccelerationStructureAddressResolutionFlagBitsNV;
```

## [](#_description)Description

- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_NONE_NV` specifies that no address resolution flags are provided.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_DST_IMPLICIT_DATA_BIT_NV` specifies another level of indirection when reading [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstImplicitData`.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_SCRATCH_DATA_BIT_NV` specifies another level of indirection when reading [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`scratchData`.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_DST_ADDRESS_ARRAY_BIT_NV` specifies another level of indirection when reading [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstAddressesArray`.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_DST_SIZES_ARRAY_BIT_NV` specifies another level of indirection when reading [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`dstSizesArray`.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_SRC_INFOS_ARRAY_BIT_NV` specifies another level of indirection when reading [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`srcInfosArray`.
- `VK_CLUSTER_ACCELERATION_STRUCTURE_ADDRESS_RESOLUTION_INDIRECTED_SRC_INFOS_COUNT_BIT_NV` specifies another level of indirection when reading [VkClusterAccelerationStructureCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureCommandsInfoNV.html)::`srcInfosCount`.

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkClusterAccelerationStructureAddressResolutionFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureAddressResolutionFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureAddressResolutionFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0