# VK_EXT_depth_range_unrestricted(3) Manual Page

## Name

VK_EXT_depth_range_unrestricted - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

14

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_range_unrestricted%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_range_unrestricted%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-06-22

**Contributors**  
- Daniel Koch, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension removes the [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) `minDepth` and
`maxDepth` restrictions that the values must be between `0.0` and `1.0`,
inclusive. It also removes the same restriction on
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
`minDepthBounds` and `maxDepthBounds`. Finally it removes the
restriction on the `depth` value in
[VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearDepthStencilValue.html).

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DEPTH_RANGE_UNRESTRICTED_EXTENSION_NAME`

- `VK_EXT_DEPTH_RANGE_UNRESTRICTED_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

1\) How do [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) `minDepth` and `maxDepth`
values outside of the `0.0` to `1.0` range interact with <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-clipping"
target="_blank" rel="noopener">Primitive Clipping</a>?

**RESOLVED**: The behavior described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-clipping"
target="_blank" rel="noopener">Primitive Clipping</a> still applies. If
depth clamping is disabled the depth values are still clipped to 0 ≤
z<sub>c</sub> ≤ w<sub>c</sub> before the viewport transform. If depth
clamping is enabled the above equation is ignored and the depth values
are instead clamped to the [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) `minDepth` and
`maxDepth` values, which in the case of this extension can be outside of
the `0.0` to `1.0` range.

2\) What happens if a resulting depth fragment is outside of the `0.0`
to `1.0` range and the depth buffer is fixed-point rather than
floating-point?

**RESOLVED**: This situation can also arise without this extension (when
fragment shaders replace depth values, for example), and this extension
does not change the behavior, which is defined in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
target="_blank" rel="noopener">Depth Test</a> section of the Fragment
Operations chapter.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-06-22 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_depth_range_unrestricted"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
