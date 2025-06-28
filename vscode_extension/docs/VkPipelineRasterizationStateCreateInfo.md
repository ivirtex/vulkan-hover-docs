# VkPipelineRasterizationStateCreateInfo(3) Manual Page

## Name

VkPipelineRasterizationStateCreateInfo - Structure specifying parameters of a newly created pipeline rasterization state



## [](#_c_specification)C Specification

The `VkPipelineRasterizationStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineRasterizationStateCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    VkPipelineRasterizationStateCreateFlags    flags;
    VkBool32                                   depthClampEnable;
    VkBool32                                   rasterizerDiscardEnable;
    VkPolygonMode                              polygonMode;
    VkCullModeFlags                            cullMode;
    VkFrontFace                                frontFace;
    VkBool32                                   depthBiasEnable;
    float                                      depthBiasConstantFactor;
    float                                      depthBiasClamp;
    float                                      depthBiasSlopeFactor;
    float                                      lineWidth;
} VkPipelineRasterizationStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `depthClampEnable` controls whether to clamp the fragment’s depth values as described in [Depth Test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth). If the pipeline is not created with [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html) present then enabling depth clamp will also disable clipping primitives to the z planes of the frustum as described in [Primitive Clipping](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-clipping). Otherwise depth clipping is controlled by the state set in [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html).
- `rasterizerDiscardEnable` controls whether primitives are discarded immediately before the rasterization stage.
- `polygonMode` is the triangle rendering mode. See [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html).
- `cullMode` is the triangle facing direction used for primitive culling. See [VkCullModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCullModeFlagBits.html).
- `frontFace` is a [VkFrontFace](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrontFace.html) value specifying the front-facing triangle orientation to be used for culling.
- `depthBiasEnable` controls whether to bias fragment depth values.
- `depthBiasConstantFactor` is a scalar factor controlling the constant depth value added to each fragment.
- `depthBiasClamp` is the maximum (or minimum) depth bias of a fragment.
- `depthBiasSlopeFactor` is a scalar factor applied to a fragment’s slope in depth bias calculations.
- `lineWidth` is the width of rasterized line segments.

## [](#_description)Description

The application **can** also add a `VkPipelineRasterizationStateRasterizationOrderAMD` structure to the `pNext` chain of a [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html) structure. This structure enables selecting the rasterization order to use when rendering with the corresponding graphics pipeline as described in [Rasterization Order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-order).

Valid Usage

- [](#VUID-VkPipelineRasterizationStateCreateInfo-depthClampEnable-00782)VUID-VkPipelineRasterizationStateCreateInfo-depthClampEnable-00782  
  If the [`depthClamp`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthClamp) feature is not enabled, `depthClampEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01507)VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01507  
  If the [`fillModeNonSolid`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fillModeNonSolid) feature is not enabled, `polygonMode` **must** be `VK_POLYGON_MODE_FILL` or `VK_POLYGON_MODE_FILL_RECTANGLE_NV`
- [](#VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01414)VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01414  
  If the `VK_NV_fill_rectangle` extension is not enabled, `polygonMode` **must** not be `VK_POLYGON_MODE_FILL_RECTANGLE_NV`
- [](#VUID-VkPipelineRasterizationStateCreateInfo-pointPolygons-04458)VUID-VkPipelineRasterizationStateCreateInfo-pointPolygons-04458  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`pointPolygons` is `VK_FALSE`, and `rasterizerDiscardEnable` is `VK_FALSE`, `polygonMode` **must** not be `VK_POLYGON_MODE_POINT`

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationStateCreateInfo-sType-sType)VUID-VkPipelineRasterizationStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO`
- [](#VUID-VkPipelineRasterizationStateCreateInfo-pNext-pNext)VUID-VkPipelineRasterizationStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthBiasRepresentationInfoEXT.html), [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html), [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html), [VkPipelineRasterizationLineStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfo.html), [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html), [VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html), or [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)
- [](#VUID-VkPipelineRasterizationStateCreateInfo-sType-unique)VUID-VkPipelineRasterizationStateCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineRasterizationStateCreateInfo-flags-zerobitmask)VUID-VkPipelineRasterizationStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-parameter)VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-parameter  
  `polygonMode` **must** be a valid [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html) value
- [](#VUID-VkPipelineRasterizationStateCreateInfo-cullMode-parameter)VUID-VkPipelineRasterizationStateCreateInfo-cullMode-parameter  
  `cullMode` **must** be a valid combination of [VkCullModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCullModeFlagBits.html) values
- [](#VUID-VkPipelineRasterizationStateCreateInfo-frontFace-parameter)VUID-VkPipelineRasterizationStateCreateInfo-frontFace-parameter  
  `frontFace` **must** be a valid [VkFrontFace](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrontFace.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCullModeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCullModeFlags.html), [VkFrontFace](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrontFace.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineRasterizationStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateFlags.html), [VkPolygonMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPolygonMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationStateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0