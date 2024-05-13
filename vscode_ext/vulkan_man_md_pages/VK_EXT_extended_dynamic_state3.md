# VK_EXT_extended_dynamic_state3(3) Manual Page

## Name

VK_EXT_extended_dynamic_state3 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

456

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

- Interacts with VK_EXT_blend_operation_advanced

- Interacts with VK_EXT_conservative_rasterization

- Interacts with VK_EXT_depth_clip_control

- Interacts with VK_EXT_depth_clip_enable

- Interacts with VK_EXT_line_rasterization

- Interacts with VK_EXT_provoking_vertex

- Interacts with VK_EXT_sample_locations

- Interacts with VK_EXT_transform_feedback

- Interacts with VK_KHR_maintenance2

- Interacts with VK_NV_clip_space_w_scaling

- Interacts with VK_NV_coverage_reduction_mode

- Interacts with VK_NV_fragment_coverage_to_color

- Interacts with VK_NV_framebuffer_mixed_samples

- Interacts with VK_NV_representative_fragment_test

- Interacts with VK_NV_shading_rate_image

- Interacts with VK_NV_viewport_swizzle

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_extended_dynamic_state3%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_extended_dynamic_state3%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_extended_dynamic_state3](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_extended_dynamic_state3.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-09-02

**IP Status**  
No known IP claims.

**Contributors**  
- Daniel Story, Nintendo

- Jamie Madill, Google

- Jan-Harald Fredriksen, Arm

- Faith Ekstrand, Collabora

- Mike Blumenkrantz, Valve

- Ricardo Garcia, Igalia

- Samuel Pitoiset, Valve

- Shahbaz Youssefi, Google

- Stu Smith, AMD

- Tapani Pälli, Intel

## <a href="#_description" class="anchor"></a>Description

This extension adds almost all of the remaining pipeline state as
dynamic state to help applications further reduce the number of
monolithic pipelines they need to create and bind.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)

- [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToOneEnableEXT.html)

- [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html)

- [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)

- [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html)

- [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClampEnableEXT.html)

- [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html)

- [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html)

- [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)

- [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html)

If
[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html)
is supported:

- [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)

If
[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html)
is supported:

- [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)

- [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)

If [VK_EXT_depth_clip_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_control.html) is
supported:

- [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)

If [VK_EXT_depth_clip_enable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_enable.html) is
supported:

- [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipEnableEXT.html)

If [VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html) is
supported:

- [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)

- [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)

If [VK_EXT_provoking_vertex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_provoking_vertex.html) is supported:

- [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetProvokingVertexModeEXT.html)

If [VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html) is supported:

- [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html)

If [VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html) is
supported:

- [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationStreamEXT.html)

If [VK_KHR_maintenance2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance2.html) or [Version
1.1](#versions-1.1) is supported:

- [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetTessellationDomainOriginEXT.html)

If [VK_NV_clip_space_w_scaling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_clip_space_w_scaling.html) is
supported:

- [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)

If [VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html)
is supported:

- [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageReductionModeNV.html)

If
[VK_NV_fragment_coverage_to_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_coverage_to_color.html)
is supported:

- [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)

- [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)

If
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html)
is supported:

- [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)

- [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)

- [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableNV.html)

If
[VK_NV_representative_fragment_test](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_representative_fragment_test.html)
is supported:

- [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)

If [VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html) is
supported:

- [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)

If [VK_NV_viewport_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_viewport_swizzle.html) is supported:

- [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendAdvancedEXT.html)

- [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorBlendEquationEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExtendedDynamicState3FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicState3FeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceExtendedDynamicState3PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicState3PropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_EXTENDED_DYNAMIC_STATE_3_EXTENSION_NAME`

- `VK_EXT_EXTENDED_DYNAMIC_STATE_3_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`

  - `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`

  - `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT`

  - `VK_DYNAMIC_STATE_POLYGON_MODE_EXT`

  - `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`

  - `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_PROPERTIES_EXT`

If
[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html)
is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`

If
[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html)
is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT`

  - `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT`

If [VK_EXT_depth_clip_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_control.html) is
supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT`

If [VK_EXT_depth_clip_enable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_enable.html) is
supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT`

If [VK_EXT_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_line_rasterization.html) is
supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT`

  - `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT`

If [VK_EXT_provoking_vertex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_provoking_vertex.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT`

If [VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT`

If [VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html) is
supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT`

If [VK_KHR_maintenance2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance2.html) or [Version
1.1](#versions-1.1) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT`

If [VK_NV_clip_space_w_scaling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_clip_space_w_scaling.html) is
supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV`

If [VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html)
is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV`

If
[VK_NV_fragment_coverage_to_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_coverage_to_color.html)
is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV`

  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV`

If
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html)
is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV`

  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV`

  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV`

If
[VK_NV_representative_fragment_test](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_representative_fragment_test.html)
is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV`

If [VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html) is
supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV`

If [VK_NV_viewport_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_viewport_swizzle.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV`

## <a href="#_issues" class="anchor"></a>Issues

1\) What about the VkPipelineMultisampleStateCreateInfo state
`sampleShadingEnable` and `minSampleShading`?

**UNRESOLVED**  
- `sampleShadingEnable` and `minSampleShading` are required when
  compiling the fragment shader, and it is not meaningful to set them
  dynamically since they always need to match the fragment shader state,
  so this hardware state may as well just come from the pipeline with
  the fragment shader.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2022-07-18 (Piers Daniell)

  - Added rasterizationSamples

- Revision 1, 2022-05-18 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_extended_dynamic_state3"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
