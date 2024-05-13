# VK_EXT_primitives_generated_query(3) Manual Page

## Name

VK_EXT_primitives_generated_query - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

383

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_primitives_generated_query%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_primitives_generated_query%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_primitives_generated_query](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_primitives_generated_query.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-01-24

**Contributors**  
- Shahbaz Youssefi, Google

- Piers Daniell, NVIDIA

- Faith Ekstrand, Collabora

- Jan-Harald Fredriksen, Arm

## <a href="#_description" class="anchor"></a>Description

This extension adds support for a new query type to match OpenGL’s
`GL_PRIMITIVES_GENERATED` to support layering.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PRIMITIVES_GENERATED_QUERY_EXTENSION_NAME`

- `VK_EXT_PRIMITIVES_GENERATED_QUERY_SPEC_VERSION`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_PRIMITIVES_GENERATED_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIMITIVES_GENERATED_QUERY_FEATURES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Can the query from `VK_EXT_transform_feedback` be used instead?

**RESOLVED**: No. While the query from VK_EXT_transform_feedback can
produce the same results as in this extension, it is only available
while transform feedback is active. The OpenGL `GL_PRIMITIVES_GENERATED`
query is independent from transform feedback. Emulation through
artificial transform feedback is unnecessarily inefficient.

2\) Can `VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT` be used
instead?

**RESOLVED**: It could, but we prefer the extension for simplicity.
Vulkan requires that only one query be active at a time. If both the
`GL_PRIMITIVES_GENERATED` and the `GL_CLIPPING_INPUT_PRIMITIVES_ARB`
queries need to be simultaneously enabled, emulation of both through
`VK_QUERY_PIPELINE_STATISTIC_CLIPPING_INVOCATIONS_BIT` is inconvenient.

3\) On some hardware, this query cannot be implemented if
`VkPipelineRasterizationStateCreateInfo`::`rasterizerDiscardEnable` is
enabled. How will this be handled?

**RESOLVED**: A feature flag is exposed by this extension for this. On
said hardware, the GL implementation disables rasterizer-discard and
achieves the same effect through other means. It will not be able to do
the same in Vulkan due to lack of state information. A feature flag is
exposed by this extension so the OpenGL implementation on top of Vulkan
would be able to implement a similar workaround.

4\) On some hardware, this query cannot be implemented for non-zero
query indices. How will this be handled?

**RESOLVED**: A feature flag is exposed by this extension for this. If
this feature is not present, the query from `VK_EXT_transform_feedback`
can be used to the same effect.

5\) How is the interaction of this extension with
`transformFeedbackRasterizationStreamSelect` handled?

**RESOLVED**: Disallowed for non-zero streams. In OpenGL, the
rasterization stream is always stream zero.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-06-23 (Shahbaz Youssefi)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_primitives_generated_query"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
