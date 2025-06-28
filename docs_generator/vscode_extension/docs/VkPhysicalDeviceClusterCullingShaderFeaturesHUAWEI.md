# VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI(3) Manual Page

## Name

VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI - Structure describing whether cluster culling shader is enabled



## [](#_c_specification)C Specification

The `VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI` structure is defined as:

```c++
// Provided by VK_HUAWEI_cluster_culling_shader
typedef struct VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           clustercullingShader;
    VkBool32           multiviewClusterCullingShader;
} VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`clustercullingShader` specifies whether cluster culling shader is supported.
- []()`multiviewClusterCullingShader` specifies whether multiview is supported.

## [](#_description)Description

If the `VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI-sType-sType)VUID-VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CLUSTER_CULLING_SHADER_FEATURES_HUAWEI`

## [](#_see_also)See Also

[VK\_HUAWEI\_cluster\_culling\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_cluster_culling_shader.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0