# VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI - Structure
describing whether cluster culling shader support VRS



## <a href="#_c_specification" class="anchor"></a>C Specification

To query whether a Cluster Culling Shader supports the per-cluster
shading rate feature, include a
`VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI` structure in the
`pNext` chain of the
`VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI` structure. This
structure is defined as:

``` c
// Provided by VK_HUAWEI_cluster_culling_shader
typedef struct VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           clusterShadingRate;
} VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-clusterShadingRate"></span> `clusterShadingRate`
  specifies whether per-cluster shading rates is supported.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI-sType-sType"
  id="VUID-VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI-sType-sType"></a>
  VUID-VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_CULLING_SHADER_VRS_FEATURES_HUAWEI`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_cluster_culling_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_cluster_culling_shader.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceClusterCullingShaderVrsFeaturesHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
