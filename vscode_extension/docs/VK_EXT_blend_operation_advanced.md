# VK\_EXT\_blend\_operation\_advanced(3) Manual Page

## Name

VK\_EXT\_blend\_operation\_advanced - device extension



## [](#_registered_extension_number)Registered Extension Number

149

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_blend_operation_advanced%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_blend_operation_advanced%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-06-12

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds a number of “advanced” blending operations that **can** be used to perform new color blending operations, many of which are more complex than the standard blend modes provided by unextended Vulkan. This extension requires different styles of usage, depending on the level of hardware support and the enabled features:

- If [VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT.html)::`advancedBlendCoherentOperations` is `VK_FALSE`, the new blending operations are supported, but a memory dependency **must** separate each advanced blend operation on a given sample. `VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT` is used to synchronize reads using advanced blend operations.
- If [VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT.html)::`advancedBlendCoherentOperations` is `VK_TRUE`, advanced blend operations obey primitive order just like basic blend operations.

In unextended Vulkan, the set of blending operations is limited, and **can** be expressed very simply. The `VK_BLEND_OP_MIN` and `VK_BLEND_OP_MAX` blend operations simply compute component-wise minimums or maximums of source and destination color components. The `VK_BLEND_OP_ADD`, `VK_BLEND_OP_SUBTRACT`, and `VK_BLEND_OP_REVERSE_SUBTRACT` modes multiply the source and destination colors by source and destination factors and either add the two products together or subtract one from the other. This limited set of operations supports many common blending operations but precludes the use of more sophisticated transparency and blending operations commonly available in many dedicated imaging APIs.

This extension provides a number of new “advanced” blending operations. Unlike traditional blending operations using `VK_BLEND_OP_ADD`, these blending equations do not use source and destination factors specified by [VkBlendFactor](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendFactor.html). Instead, each blend operation specifies a complete equation based on the source and destination colors. These new blend operations are used for both RGB and alpha components; they **must** not be used to perform separate RGB and alpha blending (via different values of color and alpha [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html)).

These blending operations are performed using premultiplied colors, where RGB colors **can** be considered premultiplied or non-premultiplied by alpha, according to the `srcPremultiplied` and `dstPremultiplied` members of [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html). If a color is considered non-premultiplied, the (R,G,B) color components are multiplied by the alpha component prior to blending. For non-premultiplied color components in the range \[0,1], the corresponding premultiplied color component would have values in the range \[0 × A, 1 × A].

Many of these advanced blending equations are formulated where the result of blending source and destination colors with partial coverage have three separate contributions: from the portions covered by both the source and the destination, from the portion covered only by the source, and from the portion covered only by the destination. The blend parameter [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap` **can** be used to specify a correlation between source and destination pixel coverage. If set to `VK_BLEND_OVERLAP_CONJOINT_EXT`, the source and destination are considered to have maximal overlap, as would be the case if drawing two objects on top of each other. If set to `VK_BLEND_OVERLAP_DISJOINT_EXT`, the source and destination are considered to have minimal overlap, as would be the case when rendering a complex polygon tessellated into individual non-intersecting triangles. If set to `VK_BLEND_OVERLAP_UNCORRELATED_EXT`, the source and destination coverage are assumed to have no spatial correlation within the pixel.

In addition to the coherency issues on implementations not supporting `advancedBlendCoherentOperations`, this extension has several limitations worth noting. First, the new blend operations have a limit on the number of color attachments they **can** be used with, as indicated by [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)::`advancedBlendMaxColorAttachments`. Additionally, blending precision **may** be limited to 16-bit floating-point, which **may** result in a loss of precision and dynamic range for framebuffer formats with 32-bit floating-point components, and in a loss of precision for formats with 12- and 16-bit signed or unsigned normalized integer components.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html)
- Extending [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateInfo.html):
  
  - [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_BLEND_OPERATION_ADVANCED_EXTENSION_NAME`
- `VK_EXT_BLEND_OPERATION_ADVANCED_SPEC_VERSION`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT`
- Extending [VkBlendOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOp.html):
  
  - `VK_BLEND_OP_BLUE_EXT`
  - `VK_BLEND_OP_COLORBURN_EXT`
  - `VK_BLEND_OP_COLORDODGE_EXT`
  - `VK_BLEND_OP_CONTRAST_EXT`
  - `VK_BLEND_OP_DARKEN_EXT`
  - `VK_BLEND_OP_DIFFERENCE_EXT`
  - `VK_BLEND_OP_DST_ATOP_EXT`
  - `VK_BLEND_OP_DST_EXT`
  - `VK_BLEND_OP_DST_IN_EXT`
  - `VK_BLEND_OP_DST_OUT_EXT`
  - `VK_BLEND_OP_DST_OVER_EXT`
  - `VK_BLEND_OP_EXCLUSION_EXT`
  - `VK_BLEND_OP_GREEN_EXT`
  - `VK_BLEND_OP_HARDLIGHT_EXT`
  - `VK_BLEND_OP_HARDMIX_EXT`
  - `VK_BLEND_OP_HSL_COLOR_EXT`
  - `VK_BLEND_OP_HSL_HUE_EXT`
  - `VK_BLEND_OP_HSL_LUMINOSITY_EXT`
  - `VK_BLEND_OP_HSL_SATURATION_EXT`
  - `VK_BLEND_OP_INVERT_EXT`
  - `VK_BLEND_OP_INVERT_OVG_EXT`
  - `VK_BLEND_OP_INVERT_RGB_EXT`
  - `VK_BLEND_OP_LIGHTEN_EXT`
  - `VK_BLEND_OP_LINEARBURN_EXT`
  - `VK_BLEND_OP_LINEARDODGE_EXT`
  - `VK_BLEND_OP_LINEARLIGHT_EXT`
  - `VK_BLEND_OP_MINUS_CLAMPED_EXT`
  - `VK_BLEND_OP_MINUS_EXT`
  - `VK_BLEND_OP_MULTIPLY_EXT`
  - `VK_BLEND_OP_OVERLAY_EXT`
  - `VK_BLEND_OP_PINLIGHT_EXT`
  - `VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT`
  - `VK_BLEND_OP_PLUS_CLAMPED_EXT`
  - `VK_BLEND_OP_PLUS_DARKER_EXT`
  - `VK_BLEND_OP_PLUS_EXT`
  - `VK_BLEND_OP_RED_EXT`
  - `VK_BLEND_OP_SCREEN_EXT`
  - `VK_BLEND_OP_SOFTLIGHT_EXT`
  - `VK_BLEND_OP_SRC_ATOP_EXT`
  - `VK_BLEND_OP_SRC_EXT`
  - `VK_BLEND_OP_SRC_IN_EXT`
  - `VK_BLEND_OP_SRC_OUT_EXT`
  - `VK_BLEND_OP_SRC_OVER_EXT`
  - `VK_BLEND_OP_VIVIDLIGHT_EXT`
  - `VK_BLEND_OP_XOR_EXT`
  - `VK_BLEND_OP_ZERO_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2017-06-12 (Jeff Bolz)
  
  - Internal revisions
- Revision 2, 2017-06-12 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_blend_operation_advanced)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0