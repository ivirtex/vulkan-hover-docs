# VK_KHR_shader_draw_parameters(3) Manual Page

## Name

VK_KHR_shader_draw_parameters - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

64

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_shader_draw_parameters](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_shader_draw_parameters.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_draw_parameters%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_draw_parameters%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_shader_draw_parameters`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_draw_parameters.txt)

**Contributors**  
- Daniel Koch, NVIDIA Corporation

- Jeff Bolz, NVIDIA

- Daniel Rakos, AMD

- Jan-Harald Fredriksen, ARM

- John Kessenich, Google

- Stuart Smith, IMG

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_KHR_shader_draw_parameters`

The extension provides access to three additional built-in shader
variables in Vulkan:

- `BaseInstance`, containing the `firstInstance` parameter passed to
  drawing commands,

- `BaseVertex`, containing the `firstVertex` or `vertexOffset` parameter
  passed to drawing commands, and

- `DrawIndex`, containing the index of the draw call currently being
  processed from an indirect drawing call.

When using GLSL source-based shader languages, the following variables
from `GL_ARB_shader_draw_parameters` can map to these SPIR-V built-in
decorations:

- `in int gl_BaseInstanceARB;` → `BaseInstance`,

- `in int gl_BaseVertexARB;` → `BaseVertex`, and

- `in int gl_DrawIDARB;` → `DrawIndex`.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1.
However, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderDrawParameters"
target="_blank" rel="noopener"><code>shaderDrawParameters</code></a>
feature bit was added to distinguish whether it is actually available or
not.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_DRAW_PARAMETERS_EXTENSION_NAME`

- `VK_KHR_SHADER_DRAW_PARAMETERS_SPEC_VERSION`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-baseinstance"
  target="_blank" rel="noopener"><code>BaseInstance</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-basevertex"
  target="_blank" rel="noopener"><code>BaseVertex</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-drawindex"
  target="_blank" rel="noopener"><code>DrawIndex</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DrawParameters"
  target="_blank" rel="noopener"><code>DrawParameters</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1\) Is this the same functionality as `GL_ARB_shader_draw_parameters`?

**RESOLVED**: It is actually a superset, as it also adds in support for
arrayed drawing commands.

In GL for `GL_ARB_shader_draw_parameters`, `gl_BaseVertexARB` holds the
integer value passed to the parameter to the command that resulted in
the current shader invocation. In the case where the command has no
`baseVertex` parameter, the value of `gl_BaseVertexARB` is zero. This
means that `gl_BaseVertexARB` = `baseVertex` (for `glDrawElements`
commands with `baseVertex`) or 0. In particular there are no
`glDrawArrays` commands that take a `baseVertex` parameter.

Now in Vulkan, we have `BaseVertex` = `vertexOffset` (for indexed
drawing commands) or `firstVertex` (for arrayed drawing commands), and
so Vulkan’s version is really a superset of GL functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-05 (Daniel Koch)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_draw_parameters"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
