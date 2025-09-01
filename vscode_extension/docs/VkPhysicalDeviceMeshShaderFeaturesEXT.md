# VkPhysicalDeviceMeshShaderFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceMeshShaderFeaturesEXT - Structure describing mesh shading features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMeshShaderFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_mesh_shader
typedef struct VkPhysicalDeviceMeshShaderFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           taskShader;
    VkBool32           meshShader;
    VkBool32           multiviewMeshShader;
    VkBool32           primitiveFragmentShadingRateMeshShader;
    VkBool32           meshShaderQueries;
} VkPhysicalDeviceMeshShaderFeaturesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`taskShader` specifies whether task shaders are supported. If this feature is not enabled, the `VK_SHADER_STAGE_TASK_BIT_EXT` and `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT` enum values **must** not be used.
- []()`meshShader` specifies whether mesh shaders are supported. If this feature is not enabled, the `VK_SHADER_STAGE_MESH_BIT_EXT` and `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT` enum values **must** not be used.
- []()`multiviewMeshShader` specifies whether the implementation supports [`multiview`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiview) rendering within a render pass, with mesh shaders. If this feature is not enabled, then a pipeline compiled against a subpass with a non-zero view mask **must** not include a mesh shader.
- []()`primitiveFragmentShadingRateMeshShader` indicates that the implementation supports the [primitive fragment shading rate](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-primitive) in mesh shaders.
- []()`meshShaderQueries` indicates that the implementation supports creating query pools using the `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT` query type and statistic queries containing the `VK_QUERY_PIPELINE_STATISTIC_TASK_SHADER_INVOCATIONS_BIT_EXT` and `VK_QUERY_PIPELINE_STATISTIC_MESH_SHADER_INVOCATIONS_BIT_EXT` flags

## [](#_description)Description

If the `VkPhysicalDeviceMeshShaderFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMeshShaderFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

The corresponding features of the `VkPhysicalDeviceMeshShaderFeaturesNV` structure **must** match those in `VkPhysicalDeviceMeshShaderFeaturesEXT`.

Valid Usage

- [](#VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-multiviewMeshShader-07032)VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-multiviewMeshShader-07032  
  If `multiviewMeshShader` is enabled then `VkPhysicalDeviceMultiviewFeaturesKHR`::`multiview` **must** also be enabled
- [](#VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-primitiveFragmentShadingRateMeshShader-07033)VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-primitiveFragmentShadingRateMeshShader-07033  
  If `primitiveFragmentShadingRateMeshShader` is enabled then `VkPhysicalDeviceFragmentShadingRateFeaturesKHR`::`primitiveFragmentShadingRate` **must** also be enabled

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_mesh\_shader](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mesh_shader.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMeshShaderFeaturesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0