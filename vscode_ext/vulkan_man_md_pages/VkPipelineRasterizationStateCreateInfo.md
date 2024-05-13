# VkPipelineRasterizationStateCreateInfo(3) Manual Page

## Name

VkPipelineRasterizationStateCreateInfo - Structure specifying parameters
of a newly created pipeline rasterization state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineRasterizationStateCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `depthClampEnable` controls whether to clamp the fragment’s depth
  values as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
  target="_blank" rel="noopener">Depth Test</a>. If the pipeline is not
  created with
  [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html)
  present then enabling depth clamp will also disable clipping
  primitives to the z planes of the frustum as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-clipping"
  target="_blank" rel="noopener">Primitive Clipping</a>. Otherwise depth
  clipping is controlled by the state set in
  [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html).

- `rasterizerDiscardEnable` controls whether primitives are discarded
  immediately before the rasterization stage.

- `polygonMode` is the triangle rendering mode. See
  [VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html).

- `cullMode` is the triangle facing direction used for primitive
  culling. See [VkCullModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlagBits.html).

- `frontFace` is a [VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html) value specifying the
  front-facing triangle orientation to be used for culling.

- `depthBiasEnable` controls whether to bias fragment depth values.

- `depthBiasConstantFactor` is a scalar factor controlling the constant
  depth value added to each fragment.

- `depthBiasClamp` is the maximum (or minimum) depth bias of a fragment.

- `depthBiasSlopeFactor` is a scalar factor applied to a fragment’s
  slope in depth bias calculations.

- `lineWidth` is the width of rasterized line segments.

## <a href="#_description" class="anchor"></a>Description

The application **can** also add a
`VkPipelineRasterizationStateRasterizationOrderAMD` structure to the
`pNext` chain of a
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
structure. This structure enables selecting the rasterization order to
use when rendering with the corresponding graphics pipeline as described
in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-order"
target="_blank" rel="noopener">Rasterization Order</a>.

Valid Usage

- <a
  href="#VUID-VkPipelineRasterizationStateCreateInfo-depthClampEnable-00782"
  id="VUID-VkPipelineRasterizationStateCreateInfo-depthClampEnable-00782"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-depthClampEnable-00782  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthClamp"
  target="_blank" rel="noopener"><code>depthClamp</code></a> feature is
  not enabled, `depthClampEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01507"
  id="VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01507"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01507  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fillModeNonSolid"
  target="_blank" rel="noopener"><code>fillModeNonSolid</code></a>
  feature is not enabled, `polygonMode` **must** be
  `VK_POLYGON_MODE_FILL` or `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

- <a href="#VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01414"
  id="VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01414"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-01414  
  If the [`VK_NV_fill_rectangle`](VK_NV_fill_rectangle.html) extension
  is not enabled, `polygonMode` **must** not be
  `VK_POLYGON_MODE_FILL_RECTANGLE_NV`

- <a
  href="#VUID-VkPipelineRasterizationStateCreateInfo-pointPolygons-04458"
  id="VUID-VkPipelineRasterizationStateCreateInfo-pointPolygons-04458"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-pointPolygons-04458  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`pointPolygons`
  is `VK_FALSE`, and `rasterizerDiscardEnable` is `VK_FALSE`,
  `polygonMode` **must** not be `VK_POLYGON_MODE_POINT`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineRasterizationStateCreateInfo-sType-sType"
  id="VUID-VkPipelineRasterizationStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineRasterizationStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineRasterizationStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html),
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html),
  [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html),
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html),
  [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html),
  [VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html),
  or
  [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)

- <a href="#VUID-VkPipelineRasterizationStateCreateInfo-sType-unique"
  id="VUID-VkPipelineRasterizationStateCreateInfo-sType-unique"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPipelineRasterizationStateCreateInfo-flags-zerobitmask"
  id="VUID-VkPipelineRasterizationStateCreateInfo-flags-zerobitmask"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-parameter"
  id="VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-parameter"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-polygonMode-parameter  
  `polygonMode` **must** be a valid [VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html)
  value

- <a
  href="#VUID-VkPipelineRasterizationStateCreateInfo-cullMode-parameter"
  id="VUID-VkPipelineRasterizationStateCreateInfo-cullMode-parameter"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-cullMode-parameter  
  `cullMode` **must** be a valid combination of
  [VkCullModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlagBits.html) values

- <a
  href="#VUID-VkPipelineRasterizationStateCreateInfo-frontFace-parameter"
  id="VUID-VkPipelineRasterizationStateCreateInfo-frontFace-parameter"></a>
  VUID-VkPipelineRasterizationStateCreateInfo-frontFace-parameter  
  `frontFace` **must** be a valid [VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkCullModeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlags.html),
[VkFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrontFace.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineRasterizationStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateFlags.html),
[VkPolygonMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPolygonMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
