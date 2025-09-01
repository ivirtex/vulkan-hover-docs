# VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI - Structure describing whether cluster culling shader support VRS



## [](#_c_specification)C Specification

To query whether a Cluster Culling Shader supports the per-cluster shading rate feature, include a `VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI` structure in the `pNext` chain of the `VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI` structure. This structure is defined as:

```c++
// Provided by VK_HUAWEI_cluster_culling_shader
typedef struct VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           clusterShadingRate;
} VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`clusterShadingRate` specifies whether per-cluster shading rates is supported.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI-sType-sType)VUID-VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_CULLING_SHADER_VRS_FEATURES_HUAWEI`

## [](#_see_also)See Also

[VK\_HUAWEI\_cluster\_culling\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_cluster_culling_shader.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0