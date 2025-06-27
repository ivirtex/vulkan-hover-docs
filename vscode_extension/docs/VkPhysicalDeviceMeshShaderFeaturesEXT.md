# VkPhysicalDeviceMeshShaderFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceMeshShaderFeaturesEXT - Structure describing mesh
shading features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMeshShaderFeaturesEXT` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-taskShader"></span> `taskShader` specifies whether
  task shaders are supported. If this feature is not enabled, the
  `VK_SHADER_STAGE_TASK_BIT_EXT` and
  `VK_PIPELINE_STAGE_TASK_SHADER_BIT_EXT` enum values **must** not be
  used.

- <span id="features-meshShader"></span> `meshShader` specifies whether
  mesh shaders are supported. If this feature is not enabled, the
  `VK_SHADER_STAGE_MESH_BIT_EXT` and
  `VK_PIPELINE_STAGE_MESH_SHADER_BIT_EXT` enum values **must** not be
  used.

- <span id="features-multiview-mesh"></span> `multiviewMeshShader`
  specifies whether the implementation supports <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> rendering
  within a render pass, with mesh shaders. If this feature is not
  enabled, then a pipeline compiled against a subpass with a non-zero
  view mask **must** not include a mesh shader.

- <span id="features-primitiveFragmentShadingRate-mesh"></span>
  `primitiveFragmentShadingRateMeshShader` indicates that the
  implementation supports the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-primitive"
  target="_blank" rel="noopener">primitive fragment shading rate</a> in
  mesh shaders.

- <span id="features-meshShaderQueries"></span> `meshShaderQueries`
  indicates that the implementation supports creating query pools using
  the `VK_QUERY_TYPE_MESH_PRIMITIVES_GENERATED_EXT` query type and
  statistic queries containing the
  `VK_QUERY_PIPELINE_STATISTIC_TASK_SHADER_INVOCATIONS_BIT_EXT` and
  `VK_QUERY_PIPELINE_STATISTIC_MESH_SHADER_INVOCATIONS_BIT_EXT` flags

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMeshShaderFeaturesEXT` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMeshShaderFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

The corresponding features of the `VkPhysicalDeviceMeshShaderFeaturesNV`
structure **must** match those in
`VkPhysicalDeviceMeshShaderFeaturesEXT`.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-multiviewMeshShader-07032"
  id="VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-multiviewMeshShader-07032"></a>
  VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-multiviewMeshShader-07032  
  If `multiviewMeshShader` is enabled then
  `VkPhysicalDeviceMultiviewFeaturesKHR`::`multiview` **must** also be
  enabled

- <a
  href="#VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-primitiveFragmentShadingRateMeshShader-07033"
  id="VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-primitiveFragmentShadingRateMeshShader-07033"></a>
  VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-primitiveFragmentShadingRateMeshShader-07033  
  If `primitiveFragmentShadingRateMeshShader` is enabled then
  `VkPhysicalDeviceFragmentShadingRateFeaturesKHR`::`primitiveFragmentShadingRate`
  **must** also be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMeshShaderFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MESH_SHADER_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_mesh_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mesh_shader.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMeshShaderFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
