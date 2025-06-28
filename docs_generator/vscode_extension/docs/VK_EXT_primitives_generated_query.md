# VK\_EXT\_primitives\_generated\_query(3) Manual Page

## Name

VK\_EXT\_primitives\_generated\_query - device extension



## [](#_registered_extension_number)Registered Extension Number

383

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_primitives_generated_query%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_primitives_generated_query%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_primitives\_generated\_query](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_primitives_generated_query.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-01-24

**Contributors**

- Shahbaz Youssefi, Google
- Piers Daniell, NVIDIA
- Faith Ekstrand, Collabora
- Jan-Harald Fredriksen, Arm

## [](#_description)Description

This extension adds support for a new query type to match OpenGL’s `GL_PRIMITIVES_GENERATED` to support layering.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PRIMITIVES_GENERATED_QUERY_EXTENSION_NAME`
- `VK_EXT_PRIMITIVES_GENERATED_QUERY_SPEC_VERSION`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIMITIVES_GENERATED_QUERY_FEATURES_EXT`

## [](#_issues)Issues

1\) Can the query from `VK_EXT_transform_feedback` be used instead?

**RESOLVED**: No. While the query from VK\_EXT\_transform\_feedback can produce the same results as in this extension, it is only available while transform feedback is active. The OpenGL `GL_PRIMITIVES_GENERATED` query is independent from transform feedback. Emulation through artificial transform feedback is unnecessarily inefficient.

2\) Can `VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT` be used instead?

**RESOLVED**: It could, but we prefer the extension for simplicity. Vulkan requires that only one query be active at a time. If both the `GL_PRIMITIVES_GENERATED` and the `GL_CLIPPING_INPUT_PRIMITIVES_ARB` queries need to be simultaneously enabled, emulation of both through `VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT` is inconvenient.

3\) On some hardware, this query cannot be implemented if `VkPipelineRasterizationStateCreateInfo`::`rasterizerDiscardEnable` is enabled. How will this be handled?

**RESOLVED**: A feature flag is exposed by this extension for this. On said hardware, the GL implementation disables rasterizer-discard and achieves the same effect through other means. It will not be able to do the same in Vulkan due to lack of state information. A feature flag is exposed by this extension so the OpenGL implementation on top of Vulkan would be able to implement a similar workaround.

4\) On some hardware, this query cannot be implemented for non-zero query indices. How will this be handled?

**RESOLVED**: A feature flag is exposed by this extension for this. If this feature is not present, the query from `VK_EXT_transform_feedback` can be used to the same effect.

5\) How is the interaction of this extension with `transformFeedbackRasterizationStreamSelect` handled?

**RESOLVED**: Disallowed for non-zero streams. In OpenGL, the rasterization stream is always stream zero.

## [](#_version_history)Version History

- Revision 1, 2021-06-23 (Shahbaz Youssefi)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_primitives_generated_query)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0