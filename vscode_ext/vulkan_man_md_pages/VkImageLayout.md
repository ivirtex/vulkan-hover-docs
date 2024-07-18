# VkImageLayout(3) Manual Page

## Name

VkImageLayout - Layout of image and image subresources



## <a href="#_c_specification" class="anchor"></a>C Specification

The set of image layouts consists of:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkImageLayout {
    VK_IMAGE_LAYOUT_UNDEFINED = 0,
    VK_IMAGE_LAYOUT_GENERAL = 1,
    VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL = 2,
    VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL = 3,
    VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL = 4,
    VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL = 5,
    VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL = 6,
    VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL = 7,
    VK_IMAGE_LAYOUT_PREINITIALIZED = 8,
  // Provided by VK_VERSION_1_1
    VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL = 1000117000,
  // Provided by VK_VERSION_1_1
    VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL = 1000117001,
  // Provided by VK_VERSION_1_2
    VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL = 1000241000,
  // Provided by VK_VERSION_1_2
    VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL = 1000241001,
  // Provided by VK_VERSION_1_2
    VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL = 1000241002,
  // Provided by VK_VERSION_1_2
    VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL = 1000241003,
  // Provided by VK_VERSION_1_3
    VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL = 1000314000,
  // Provided by VK_VERSION_1_3
    VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL = 1000314001,
  // Provided by VK_KHR_swapchain
    VK_IMAGE_LAYOUT_PRESENT_SRC_KHR = 1000001002,
  // Provided by VK_KHR_video_decode_queue
    VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR = 1000024000,
  // Provided by VK_KHR_video_decode_queue
    VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR = 1000024001,
  // Provided by VK_KHR_video_decode_queue
    VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR = 1000024002,
  // Provided by VK_KHR_shared_presentable_image
    VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR = 1000111000,
  // Provided by VK_EXT_fragment_density_map
    VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT = 1000218000,
  // Provided by VK_KHR_fragment_shading_rate
    VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR = 1000164003,
  // Provided by VK_KHR_dynamic_rendering_local_read
    VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR = 1000232000,
  // Provided by VK_KHR_video_encode_queue
    VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR = 1000299000,
  // Provided by VK_KHR_video_encode_queue
    VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR = 1000299001,
  // Provided by VK_KHR_video_encode_queue
    VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR = 1000299002,
  // Provided by VK_EXT_attachment_feedback_loop_layout
    VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT = 1000339000,
  // Provided by VK_KHR_maintenance2
    VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR = VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL,
  // Provided by VK_KHR_maintenance2
    VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR = VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL,
  // Provided by VK_NV_shading_rate_image
    VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV = VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR,
  // Provided by VK_KHR_separate_depth_stencil_layouts
    VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR = VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL,
  // Provided by VK_KHR_separate_depth_stencil_layouts
    VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR = VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL,
  // Provided by VK_KHR_separate_depth_stencil_layouts
    VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL_KHR = VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL,
  // Provided by VK_KHR_separate_depth_stencil_layouts
    VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR = VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL,
  // Provided by VK_KHR_synchronization2
    VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR = VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL,
  // Provided by VK_KHR_synchronization2
    VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR = VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL,
} VkImageLayout;
```

## <a href="#_description" class="anchor"></a>Description

The type(s) of device access supported by each layout are:

- `VK_IMAGE_LAYOUT_UNDEFINED` specifies that the layout is unknown.
  Image memory **cannot** be transitioned into this layout. This layout
  **can** be used as the `initialLayout` member of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html). This layout **can** be
  used in place of the current image layout in a layout transition, but
  doing so will cause the contents of the image’s memory to be
  undefined.

- `VK_IMAGE_LAYOUT_PREINITIALIZED` specifies that an image’s memory is
  in a defined layout and **can** be populated by data, but that it has
  not yet been initialized by the driver. Image memory **cannot** be
  transitioned into this layout. This layout **can** be used as the
  `initialLayout` member of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html).
  This layout is intended to be used as the initial layout for an image
  whose contents are written by the host, and hence the data **can** be
  written to memory immediately, without first executing a layout
  transition. Currently, `VK_IMAGE_LAYOUT_PREINITIALIZED` is only useful
  with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-linear-resource"
  target="_blank" rel="noopener">linear</a> images because there is not
  a standard layout defined for `VK_IMAGE_TILING_OPTIMAL` images.

- `VK_IMAGE_LAYOUT_GENERAL` supports all types of device access.

- `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL` specifies a layout that **must**
  only be used with attachment accesses in the graphics pipeline.

- `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL` specifies a layout allowing read
  only access as an attachment, or in shaders as a sampled image,
  combined image/sampler, or input attachment.

- `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` **must** only be used as a
  color or resolve attachment in a `VkFramebuffer`. This layout is valid
  only for image subresources of images created with the
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` usage bit enabled.

- `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` specifies a layout
  for both the depth and stencil aspects of a depth/stencil format image
  allowing read and write access as a depth/stencil attachment. It is
  equivalent to `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` and
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`.

- `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL` specifies a layout
  for both the depth and stencil aspects of a depth/stencil format image
  allowing read only access as a depth/stencil attachment or in shaders
  as a sampled image, combined image/sampler, or input attachment. It is
  equivalent to `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` and
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`.

- `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` specifies
  a layout for depth/stencil format images allowing read and write
  access to the stencil aspect as a stencil attachment, and read only
  access to the depth aspect as a depth attachment or in shaders as a
  sampled image, combined image/sampler, or input attachment. It is
  equivalent to `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` and
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`.

- `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` specifies
  a layout for depth/stencil format images allowing read and write
  access to the depth aspect as a depth attachment, and read only access
  to the stencil aspect as a stencil attachment or in shaders as a
  sampled image, combined image/sampler, or input attachment. It is
  equivalent to `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` and
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`.

- `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` specifies a layout for the
  depth aspect of a depth/stencil format image allowing read and write
  access as a depth attachment.

- `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` specifies a layout for the
  depth aspect of a depth/stencil format image allowing read-only access
  as a depth attachment or in shaders as a sampled image, combined
  image/sampler, or input attachment.

- `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` specifies a layout for
  the stencil aspect of a depth/stencil format image allowing read and
  write access as a stencil attachment.

- `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` specifies a layout for the
  stencil aspect of a depth/stencil format image allowing read-only
  access as a stencil attachment or in shaders as a sampled image,
  combined image/sampler, or input attachment.

- `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` specifies a layout allowing
  read-only access in a shader as a sampled image, combined
  image/sampler, or input attachment. This layout is valid only for
  image subresources of images created with the
  `VK_IMAGE_USAGE_SAMPLED_BIT` or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
  usage bits enabled.

- `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` **must** only be used as a
  source image of a transfer command (see the definition of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-transfer"
  target="_blank"
  rel="noopener"><code>VK_PIPELINE_STAGE_TRANSFER_BIT</code></a>). This
  layout is valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` usage bit enabled.

- `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` **must** only be used as a
  destination image of a transfer command. This layout is valid only for
  image subresources of images created with the
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` usage bit enabled.

- `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR` **must** only be used for presenting
  a presentable image for display.

- `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR` is valid only for shared
  presentable images, and **must** be used for any usage the image
  supports.

- `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`
  **must** only be used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-shading-rate-image"
  target="_blank" rel="noopener">shading rate image</a>. This layout is
  valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR` usage bit
  enabled.

- `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT` **must** only be
  used as a fragment density map attachment in a `VkRenderPass`. This
  layout is valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT` usage bit enabled.

- `VK_IMAGE_LAYOUT_VIDEO_DECODE_DST_KHR` **must** only be used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
  target="_blank" rel="noopener">decode output picture</a> in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operation</a>. This layout
  is valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR` usage bit enabled.

- `VK_IMAGE_LAYOUT_VIDEO_DECODE_SRC_KHR` is reserved for future use.

- `VK_IMAGE_LAYOUT_VIDEO_DECODE_DPB_KHR` **must** only be used as an
  output <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> or an input
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference picture</a> in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operation</a>. This layout
  is valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR` usage bit enabled.

- `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DST_KHR` is reserved for future use.

- `VK_IMAGE_LAYOUT_VIDEO_ENCODE_SRC_KHR` **must** only be used as an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
  target="_blank" rel="noopener">encode input picture</a> in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operation</a>. This layout
  is valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR` usage bit enabled.

- `VK_IMAGE_LAYOUT_VIDEO_ENCODE_DPB_KHR` **must** only be used as an
  output <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reconstructed-picture"
  target="_blank" rel="noopener">reconstructed picture</a> or an input
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference picture</a> in a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operation</a>. This layout
  is valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR` usage bit enabled.

- `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` **must** only
  be used as either a color attachment or depth/stencil attachment in a
  `VkFramebuffer` and/or read-only access in a shader as a sampled
  image, combined image/sampler, or input attachment. This layout is
  valid only for image subresources of images created with the
  `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` usage bit enabled
  and either the `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` and either the
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT`
  usage bits enabled.

- `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR` **must** only be used as
  either a storage image, or a color or depth/stencil attachment and an
  input attachment. This layout is valid only for image subresources of
  images created with either `VK_IMAGE_USAGE_STORAGE_BIT`, or both
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` and either of
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`.

The layout of each image subresource is not a state of the image
subresource itself, but is rather a property of how the data in memory
is organized, and thus for each mechanism of accessing an image in the
API the application **must** specify a parameter or structure member
that indicates which image layout the image subresource(s) are
considered to be in when the image will be accessed. For transfer
commands, this is a parameter to the command (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears"
class="bare" target="_blank"
rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears</a>
and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies"
class="bare" target="_blank"
rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#copies</a>).
For use as a framebuffer attachment, this is a member in the
substructures of the
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass"
target="_blank" rel="noopener">Render Pass</a>). For use in a descriptor
set, this is a member in the `VkDescriptorImageInfo` structure (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-updates"
class="bare" target="_blank"
rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-updates</a>).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html),
[VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html),
[VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html),
[VkAttachmentReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference.html),
[VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html),
[VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayout.html),
[VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html),
[VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferToImageInfo2.html),
[VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageInfo2.html),
[VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToBufferInfo2.html),
[VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html),
[VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html),
[VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html),
[VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html),
[VkHostImageLayoutTransitionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageLayoutTransitionInfoEXT.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html),
[VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2.html),
[VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html),
[VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html),
[VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html),
[VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html),
[VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveImageInfo2.html),
[vkBindOpticalFlowSessionImageNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindOpticalFlowSessionImageNV.html),
[vkCmdBindInvocationMaskHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindInvocationMaskHUAWEI.html),
[vkCmdBindShadingRateImageNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindShadingRateImageNV.html),
[vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage.html),
[vkCmdClearColorImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearColorImage.html),
[vkCmdClearDepthStencilImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearDepthStencilImage.html),
[vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html),
[vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage.html),
[vkCmdCopyImageToBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer.html),
[vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToImageIndirectNV.html),
[vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
