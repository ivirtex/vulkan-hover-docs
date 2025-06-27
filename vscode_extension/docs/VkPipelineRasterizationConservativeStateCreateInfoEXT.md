# VkPipelineRasterizationConservativeStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationConservativeStateCreateInfoEXT - Structure
specifying conservative raster state



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
includes a `VkPipelineRasterizationConservativeStateCreateInfoEXT`
structure, then that structure includes parameters controlling
conservative rasterization.

`VkPipelineRasterizationConservativeStateCreateInfoEXT` is defined as:

``` c
// Provided by VK_EXT_conservative_rasterization
typedef struct VkPipelineRasterizationConservativeStateCreateInfoEXT {
    VkStructureType                                           sType;
    const void*                                               pNext;
    VkPipelineRasterizationConservativeStateCreateFlagsEXT    flags;
    VkConservativeRasterizationModeEXT                        conservativeRasterizationMode;
    float                                                     extraPrimitiveOverestimationSize;
} VkPipelineRasterizationConservativeStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `conservativeRasterizationMode` is the conservative rasterization mode
  to use.

- `extraPrimitiveOverestimationSize` is the extra size in pixels to
  increase the generating primitive during conservative rasterization at
  each of its edges in `X` and `Y` equally in screen space beyond the
  base overestimation specified in
  `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`primitiveOverestimationSize`.
  If `conservativeRasterizationMode` is not
  `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT`, this value is
  ignored.

## <a href="#_description" class="anchor"></a>Description

If this structure is not included in the `pNext` chain,
`conservativeRasterizationMode` is considered to be
`VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT`, and conservative
rasterization is disabled.

Polygon rasterization **can** be made conservative by setting
`conservativeRasterizationMode` to
`VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT` or
`VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT` in
`VkPipelineRasterizationConservativeStateCreateInfoEXT`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-conservativePointAndLineRasterization"
target="_blank"
rel="noopener"><code>conservativePointAndLineRasterization</code></a> is
supported, conservative rasterization can be applied to line and point
primitives, otherwise it must be disabled.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-extraPrimitiveOverestimationSize-01769"
  id="VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-extraPrimitiveOverestimationSize-01769"></a>
  VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-extraPrimitiveOverestimationSize-01769  
  `extraPrimitiveOverestimationSize` **must** be in the range of `0.0`
  to
  `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`maxExtraPrimitiveOverestimationSize`
  inclusive

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-conservativeRasterizationMode-parameter"
  id="VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-conservativeRasterizationMode-parameter"></a>
  VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-conservativeRasterizationMode-parameter  
  `conservativeRasterizationMode` **must** be a valid
  [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConservativeRasterizationModeEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html),
[VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConservativeRasterizationModeEXT.html),
[VkPipelineRasterizationConservativeStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationConservativeStateCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
