# VK\_EXT\_extended\_dynamic\_state3(3) Manual Page

## Name

VK\_EXT\_extended\_dynamic\_state3 - device extension



## [](#_registered_extension_number)Registered Extension Number

456

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_1
- Interacts with VK\_EXT\_blend\_operation\_advanced
- Interacts with VK\_EXT\_conservative\_rasterization
- Interacts with VK\_EXT\_depth\_clip\_control
- Interacts with VK\_EXT\_depth\_clip\_enable
- Interacts with VK\_EXT\_line\_rasterization
- Interacts with VK\_EXT\_provoking\_vertex
- Interacts with VK\_EXT\_sample\_locations
- Interacts with VK\_EXT\_transform\_feedback
- Interacts with VK\_KHR\_maintenance2
- Interacts with VK\_NV\_clip\_space\_w\_scaling
- Interacts with VK\_NV\_coverage\_reduction\_mode
- Interacts with VK\_NV\_fragment\_coverage\_to\_color
- Interacts with VK\_NV\_framebuffer\_mixed\_samples
- Interacts with VK\_NV\_representative\_fragment\_test
- Interacts with VK\_NV\_shading\_rate\_image
- Interacts with VK\_NV\_viewport\_swizzle
- Interacts with VkPhysicalDeviceExtendedDynamicState3FeaturesEXT::extendedDynamicState3AlphaToOneEnable
- Interacts with VkPhysicalDeviceExtendedDynamicState3FeaturesEXT::extendedDynamicState3DepthClampEnable
- Interacts with VkPhysicalDeviceExtendedDynamicState3FeaturesEXT::extendedDynamicState3LogicOpEnable
- Interacts with VkPhysicalDeviceExtendedDynamicState3FeaturesEXT::extendedDynamicState3PolygonMode
- Interacts with VkPhysicalDeviceExtendedDynamicState3FeaturesEXT::extendedDynamicState3RasterizationStream
- Interacts with VkPhysicalDeviceExtendedDynamicState3FeaturesEXT::extendedDynamicState3TessellationDomainOrigin

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_extended_dynamic_state3%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_extended_dynamic_state3%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_extended\_dynamic\_state3](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_extended_dynamic_state3.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

This extension adds almost all of the remaining pipeline state as dynamic state to help applications further reduce the number of monolithic pipelines they need to create and bind.

## [](#_new_commands)New Commands

- [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)
- [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetAlphaToOneEnableEXT.html)
- [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendEnableEXT.html)
- [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendEquationEXT.html)
- [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorWriteMaskEXT.html)
- [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClampEnableEXT.html)
- [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLogicOpEnableEXT.html)
- [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPolygonModeEXT.html)
- [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizationSamplesEXT.html)
- [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetSampleMaskEXT.html)

If [VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html) is supported:

- [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetColorBlendAdvancedEXT.html)

If [VK\_EXT\_conservative\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conservative_rasterization.html) is supported:

- [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
- [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)

If [VK\_EXT\_depth\_clip\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_control.html) is supported:

- [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)

If [VK\_EXT\_depth\_clip\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_enable.html) is supported:

- [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClipEnableEXT.html)

If [VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html) is supported:

- [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineRasterizationModeEXT.html)
- [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineStippleEnableEXT.html)

If [VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html) is supported:

- [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetProvokingVertexModeEXT.html)

If [VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html) is supported:

- [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetSampleLocationsEnableEXT.html)

If [VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html) is supported:

- [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRasterizationStreamEXT.html)

If [VK\_KHR\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance2.html) or [Vulkan Version 1.1](#versions-1.1) is supported:

- [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetTessellationDomainOriginEXT.html)

If [VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html) is supported:

- [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWScalingEnableNV.html)

If [VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html) is supported:

- [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageReductionModeNV.html)

If [VK\_NV\_fragment\_coverage\_to\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_coverage_to_color.html) is supported:

- [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageToColorEnableNV.html)
- [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageToColorLocationNV.html)

If [VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html) is supported:

- [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationModeNV.html)
- [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
- [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationTableNV.html)

If [VK\_NV\_representative\_fragment\_test](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_representative_fragment_test.html) is supported:

- [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)

If [VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html) is supported:

- [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetShadingRateImageEnableNV.html)

If [VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html) is supported:

- [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportSwizzleNV.html)

## [](#_new_structures)New Structures

- [VkColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendAdvancedEXT.html)
- [VkColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkColorBlendEquationEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExtendedDynamicState3FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedDynamicState3FeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceExtendedDynamicState3PropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedDynamicState3PropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTENDED_DYNAMIC_STATE_3_EXTENSION_NAME`
- `VK_EXT_EXTENDED_DYNAMIC_STATE_3_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
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
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_3_PROPERTIES_EXT`

If [VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`

If [VK\_EXT\_conservative\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conservative_rasterization.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT`
  - `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT`

If [VK\_EXT\_depth\_clip\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_control.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT`

If [VK\_EXT\_depth\_clip\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_enable.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT`

If [VK\_EXT\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_line_rasterization.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT`
  - `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT`

If [VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT`

If [VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT`

If [VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT`

If [VK\_KHR\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance2.html) or [Vulkan Version 1.1](#versions-1.1) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT`

If [VK\_NV\_clip\_space\_w\_scaling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_clip_space_w_scaling.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV`

If [VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV`

If [VK\_NV\_fragment\_coverage\_to\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_coverage_to_color.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV`
  - `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV`

If [VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV`
  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV`
  - `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV`

If [VK\_NV\_representative\_fragment\_test](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_representative_fragment_test.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV`

If [VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV`

If [VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html) is supported:

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV`

## [](#_issues)Issues

1\) What about the VkPipelineMultisampleStateCreateInfo state `sampleShadingEnable` and `minSampleShading`?

**UNRESOLVED**

- `sampleShadingEnable` and `minSampleShading` are required when compiling the fragment shader, and it is not meaningful to set them dynamically since they always need to match the fragment shader state, so this hardware state may as well just come from the pipeline with the fragment shader.

## [](#_version_history)Version History

- Revision 2, 2022-07-18 (Piers Daniell)
  
  - Added rasterizationSamples
- Revision 1, 2022-05-18 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_extended_dynamic_state3)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0