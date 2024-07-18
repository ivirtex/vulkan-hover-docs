# VK_AMD_rasterization_order(3) Manual Page

## Name

VK_AMD_rasterization_order - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

19

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_rasterization_order%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_rasterization_order%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-04-25

**IP Status**  
No known IP claims.

**Contributors**  
- Matthaeus G. Chajdas, AMD

- Jaakko Konttinen, AMD

- Daniel Rakos, AMD

- Graham Sellers, AMD

- Dominik Witczak, AMD

## <a href="#_description" class="anchor"></a>Description

This extension introduces the possibility for the application to control
the order of primitive rasterization. In unextended Vulkan, the
following stages are guaranteed to execute in *API order*:

- depth bounds test

- stencil test, stencil op, and stencil write

- depth test and depth write

- occlusion queries

- blending, logic op, and color write

This extension enables applications to opt into a relaxed,
implementation defined primitive rasterization order that may allow
better parallel processing of primitives and thus enabling higher
primitive throughput. It is applicable in cases where the primitive
rasterization order is known to not affect the output of the rendering
or any differences caused by a different rasterization order are not a
concern from the point of view of the application’s purpose.

A few examples of cases when using the relaxed primitive rasterization
order would not have an effect on the final rendering:

- If the primitives rendered are known to not overlap in framebuffer
  space.

- If depth testing is used with a comparison operator of
  `VK_COMPARE_OP_LESS`, `VK_COMPARE_OP_LESS_OR_EQUAL`,
  `VK_COMPARE_OP_GREATER`, or `VK_COMPARE_OP_GREATER_OR_EQUAL`, and the
  primitives rendered are known to not overlap in clip space.

- If depth testing is not used and blending is enabled for all
  attachments with a commutative blend operator.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html):

  - [VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRasterizationOrderAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_RASTERIZATION_ORDER_EXTENSION_NAME`

- `VK_AMD_RASTERIZATION_ORDER_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD`

## <a href="#_issues" class="anchor"></a>Issues

1\) How is this extension useful to application developers?

**RESOLVED**: Allows them to increase primitive throughput for cases
when strict API order rasterization is not important due to the nature
of the content, the configuration used, or the requirements towards the
output of the rendering.

2\) How does this extension interact with content optimizations aiming
to reduce overdraw by appropriately ordering the input primitives?

**RESOLVED**: While the relaxed rasterization order might somewhat limit
the effectiveness of such content optimizations, most of the benefits of
it are expected to be retained even when the relaxed rasterization order
is used, so applications **should** still apply these optimizations even
if they intend to use the extension.

3\) Are there any guarantees about the primitive rasterization order
when using the new relaxed mode?

**RESOLVED**: No. In this case the rasterization order is completely
implementation-dependent, but in practice it is expected to partially
still follow the order of incoming primitives.

4\) Does the new relaxed rasterization order have any adverse effect on
repeatability and other invariance rules of the API?

**RESOLVED**: Yes, in the sense that it extends the list of exceptions
when the repeatability requirement does not apply.

## <a href="#_examples" class="anchor"></a>Examples

None

## <a href="#_issues_2" class="anchor"></a>Issues

None

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-04-25 (Daniel Rakos)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_rasterization_order"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
