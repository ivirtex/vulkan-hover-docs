# VK\_QCOM\_tile\_shading(3) Manual Page

## Name

VK\_QCOM\_tile\_shading - device extension



## [](#_registered_extension_number)Registered Extension Number

310

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_QCOM\_tile\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_properties.html)  
or  
[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_QCOM\_tile\_shading](https://github.khronos.org/SPIRV-Registry/extensions/QCOM/SPV_QCOM_tile_shading.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_tile_shading%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_tile_shading%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_QCOM\_tile\_shading](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_QCOM_tile_shading.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-8-13

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension interacts with `VK_KHR_dynamic_rendering`
- This extension interacts with `VK_EXT_transform_feedback`
- This extension interacts with `VK_EXT_debug_marker`
- This extension interacts with `VK_EXT_attachment_feedback_loop_layout`
- This extension interacts with `VK_KHR_dynamic_rendering_local_read`
- This extension interacts with `VK_QCOM_image_processing`

**Contributors**

- Jeff Leger, Qualcomm
- Matt Netsch, Qualcomm
- Srihari Babu Alla, Qualcomm
- Matthew Smith, Qualcomm
- Kevin Matlage, Qualcomm
- Alex Bourd, Qualcomm

## [](#_description)Description

This extension exposes tile shading in Vulkan. Many mobile GPUs utilize Tile-Based Deferred Rendering (TBDR) to optimize for power and performance. Conversely, most desktop GPUs use immediate-mode rendering (IM). Adreno ™ GPUs uniquely have the ability to operate in either mode, and when this extension is not enabled, the Adreno driver will select the most optimal mode (TBDR or IM) based on the workload; this feature is called FlexRender ™. When this extension is in use, FlexRender is disabled and the GPU operates exclusively in TBDR wherever possible.

The TBDR mode divides the color and depth/stencil buffer attachments into a regular grid of smaller regions called "tiles". When a render pass instance is submitted for execution on an Adreno GPU, the rendering is split into two phases: a single "visibility pass" followed by multiple "rendering passes" where a separate render pass is issued for each tile in the framebuffer.

The "visibility pass" processes the geometry: identifies which tiles are covered by each primitive, eliminates occluded primitives and unneeded state changes, and performs other tile-specific optimizations. The primitive coverage information collected during the visibility pass is used in the subsequent "rendering pass" for each tile. During the rendering pass for each tile, any primitives that were determined not to cover the current tile are skipped.

This deferred rasterization additionally utilizes a specialized high-bandwidth on-die memory, "tile memory". Tile memory is dramatically more efficient than other device memory. The tile memory temporarily stores the color and other attachments for each tile during rasterization. After each tile is fully rasterized, the resulting tile is typically copied to device memory backing the attachment as specified by the render pass STORE\_OP. The per-tile rendering passes occur independently for each tile, with multiple tiles potentially being processed in parallel.

This extension enables applications to leverage the power and performance of tile memory in new ways:

- Adds a mechanism for recording dispatches or draws that are guaranteed to be executed per-tile.
- Such draws bypass the above-mentioned visibility-based skipping and are guaranteed to be executed for every tile in the rendering pass.
- Shaders can declare "tile attachments" variables, providing shader access to color, depth/stencil, and input attachment pixels.
- Fragment and compute shaders can read these render pass attachments at any location within the tile. Compute shaders can also write to color attachments at any location within the tile.
- Shaders can use new built-in variables that provide the location, size, and apron region of the tile.
- A new tile dispatch command automatically scales workgroup sizes and counts to the tile size, given a desired shading rate.
- Framebuffer-local dependencies are expanded to tile-sized regions, rather than a single pixel or sample.
- A tile shading render pass can also enable tiling "aprons". This is a specialized rendering mode where the GPU renders overlapping tiles that enable specific use cases.

## [](#_new_commands)New Commands

- [vkCmdBeginPerTileExecutionQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginPerTileExecutionQCOM.html)
- [vkCmdDispatchTileQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchTileQCOM.html)
- [vkCmdEndPerTileExecutionQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndPerTileExecutionQCOM.html)

## [](#_new_structures)New Structures

- [VkDispatchTileInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDispatchTileInfoQCOM.html)
- [VkPerTileBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileBeginInfoQCOM.html)
- [VkPerTileEndInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerTileEndInfoQCOM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceTileShadingFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTileShadingFeaturesQCOM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceTileShadingPropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTileShadingPropertiesQCOM.html)
- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html), [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html):
  
  - [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)

## [](#_new_enums)New Enums

- [VkTileShadingRenderPassFlagBitsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileShadingRenderPassFlagBitsQCOM.html)

## [](#_new_bitmasks)New Bitmasks

- [VkTileShadingRenderPassFlagsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileShadingRenderPassFlagsQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_TILE_SHADING_EXTENSION_NAME`
- `VK_QCOM_TILE_SHADING_SPEC_VERSION`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_READ_BIT_QCOM`
  - `VK_ACCESS_2_SHADER_TILE_ATTACHMENT_WRITE_BIT_QCOM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DISPATCH_TILE_INFO_QCOM`
  - `VK_STRUCTURE_TYPE_PER_TILE_BEGIN_INFO_QCOM`
  - `VK_STRUCTURE_TYPE_PER_TILE_END_INFO_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_SHADING_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_SHADING_PROPERTIES_QCOM`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_TILE_SHADING_CREATE_INFO_QCOM`
- Extending [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html):
  
  - `VK_SUBPASS_DESCRIPTION_TILE_SHADING_APRON_BIT_QCOM`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`TileOffsetQCOM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-tileoffset)
- [`TileDimensionQCOM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-tilesize)
- [`TileApronSizeQCOM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-tileapronsize)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`TileShadingQCOM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-TileShadingQCOM)

## [](#_issues)Issues

1\) Some early Adreno drivers advertised support for version 1 of this extension without supporting the required [`tileShadingApron`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingApron) feature. To cover all Adreno devices on the market, applications should not assume any version of this extension supports the `tileShadingApron` feature without performing a feature query.

## [](#_version_history)Version History

- Revision 2, 2025-08-13 (Matthew Netsch)
  
  - Make the `tileShadingApron` feature optional
- Revision 1, 2023-10-12 (Jeff Leger)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_tile_shading).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0