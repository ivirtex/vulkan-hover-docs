# VK_EXT_multisampled_render_to_single_sampled(3) Manual Page

## Name

VK_EXT_multisampled_render_to_single_sampled - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

377

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html)  
     and  
     [VK_KHR_depth_stencil_resolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_depth_stencil_resolve.html)  
or  
[Version 1.2](#versions-1.2)  

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_multisampled_render_to_single_sampled%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_multisampled_render_to_single_sampled%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_multisampled_render_to_single_sampled](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_multisampled_render_to_single_sampled.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-04-16

**IP Status**  
No known IP claims.

**Contributors**  
- Shahbaz Youssefi, Google

- Jan-Harald Fredriksen, Arm

- Jörg Wagner, Arm

- Matthew Netsch, Qualcomm Technologies, Inc.

- Jarred Davies, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

With careful usage of resolve attachments, multisampled image memory
allocated with `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT`, `loadOp` not
equal to `VK_ATTACHMENT_LOAD_OP_LOAD` and `storeOp` not equal to
`VK_ATTACHMENT_STORE_OP_STORE`, a Vulkan application is able to
efficiently perform multisampled rendering without incurring any
additional memory penalty on some implementations.

Under certain circumstances however, the application may not be able to
complete its multisampled rendering within a single render pass; for
example if it does partial rasterization from frame to frame, blending
on an image from a previous frame, or in emulation of
GL_EXT_multisampled_render_to_texture. In such cases, the application
can use an initial subpass to effectively load single-sampled data from
the next subpass’s resolve attachment and fill in the multisampled
attachment which otherwise uses `loadOp` equal to
`VK_ATTACHMENT_LOAD_OP_DONT_CARE`. However, this is not always possible
(for example for stencil in the absence of VK_EXT_shader_stencil_export)
and has multiple drawbacks.

Some implementations are able to perform said operation efficiently in
hardware, effectively loading a multisampled attachment from the
contents of a single sampled one. Together with the ability to perform a
resolve operation at the end of a subpass, these implementations are
able to perform multisampled rendering on single-sampled attachments
with no extra memory or bandwidth overhead. This extension exposes this
capability by allowing a framebuffer and render pass to include
single-sampled attachments while rendering is done with a specified
number of samples.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html):

  - [VkSubpassResolvePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassResolvePerformanceQueryEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT.html)

- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html),
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_EXTENSION_NAME`

- `VK_EXT_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_SPEC_VERSION`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_SUBPASS_RESOLVE_PERFORMANCE_QUERY_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Could the multisampled attachment be initialized through some form
of copy?

**RESOLVED**: No. Some implementations do not support copying between
attachments in general, and find expressing this operation through a
copy unnatural.

2\) Another way to achieve this is by introducing a new `loadOp` to load
the contents of the multisampled image from a single-sampled one. Why is
this extension preferred?

**RESOLVED**: Using this extension simplifies the application, as it
does not need to manage a secondary lazily-allocated image.
Additionally, using this extension leaves less room for error; for
example a single mistake in `loadOp` or `storeOp` would result in the
lazily-allocated image to actually take up memory, and remain so until
destruction.

3\) There is no guarantee that multisampled data between two subpasses
with the same number of samples will be retained as the implementation
may be forced to split the render pass implicitly for various reasons.
Should this extension require that every subpass that uses
multisampled-render-to-single-sampled end in an implicit render pass
split (which results in a resolve operation)?

**RESOLVED**: No. Not requiring this allows render passes with multiple
multisampled-render-to-single-sampled subpasses to potentially execute
more efficiently (though there is no guarantee).

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-04-12 (Shahbaz Youssefi)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_multisampled_render_to_single_sampled"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
