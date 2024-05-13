# VK_EXT_fragment_density_map(3) Manual Page

## Name

VK_EXT_fragment_density_map - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

219

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_format_feature_flags2

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_fragment_invocation_density](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_fragment_invocation_density.html)

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_fragment_density_map%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_fragment_density_map%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-09-30

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_fragment_invocation_density`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_fragment_invocation_density.txt)

**Contributors**  
- Matthew Netsch, Qualcomm Technologies, Inc.

- Robert VanReenen, Qualcomm Technologies, Inc.

- Jonathan Wicks, Qualcomm Technologies, Inc.

- Tate Hornbeck, Qualcomm Technologies, Inc.

- Sam Holmes, Qualcomm Technologies, Inc.

- Jeff Leger, Qualcomm Technologies, Inc.

- Jan-Harald Fredriksen, ARM

- Jeff Bolz, NVIDIA

- Pat Brown, NVIDIA

- Daniel Rakos, AMD

- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to specify areas of the render
target where the fragment shader may be invoked fewer times. These
fragments are broadcasted out to multiple pixels to cover the render
target.

The primary use of this extension is to reduce workloads in areas where
lower quality may not be perceived such as the distorted edges of a lens
or the periphery of a user’s gaze.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentDensityMapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceFragmentDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapPropertiesEXT.html)

- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html),
  [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html):

  - [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_FRAGMENT_DENSITY_MAP_EXTENSION_NAME`

- `VK_EXT_FRAGMENT_DENSITY_MAP_SPEC_VERSION`

- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html):

  - `VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`

- Extending [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlagBits.html):

  - `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT`

- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits.html):

  - `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlagBits.html):

  - `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`

  - `VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT`

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html):

  - `VK_FORMAT_FEATURE_2_FRAGMENT_DENSITY_MAP_BIT_EXT`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-fraginvocationcount"
  target="_blank" rel="noopener"><code>FragInvocationCountEXT</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-fragsize"
  target="_blank" rel="noopener"><code>FragSizeEXT</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentDensityEXT"
  target="_blank" rel="noopener"><code>FragmentDensityEXT</code></a>

## <a href="#_examples" class="anchor"></a>Examples

### <a href="#_fragment_density_map" class="anchor"></a>Fragment Density Map

An image can be bound as a fragment density map attachment to a render
pass. This image contains normalized (x, y) float component fragment
density values for regions of the framebuffer that will be used in
rasterization for every subpass. A float component ranges from (0.0,
1.0\] where 1.0 means full density along that axis. Implementations <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragmentdensitymapops"
target="_blank" rel="noopener">use these values as hints</a> to optimize
rendering in areas of low density. Subpass color and depth attachments
can be created as subsampled, which can help to further optimize
rendering in areas of low density.

The density map image can be modified by the application until calling
`vkCmdBeginRenderPass` for the render pass that uses the image. If
`VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` is used,
then the application can modify the image until the device reads it
during `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`.

``` c
// Create fragment density map
VkImageCreateInfo imageCreateInfo =
{
   .sType = VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO,
   .pNext = nullptr,
   .flags = 0,
   .imageType = VK_IMAGE_TYPE_2D,    // Must be 2D
   .format = VK_FORMAT_R8G8_UNORM,   // Must have VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT
   .extend = {64, 64, 1},
   .mipLevels = 1,
   .arrayLayers = 2,                 // 1 for each multiview view
   .samples = VK_SAMPLE_COUNT_1_BIT, // Must be 1x MSAA
   .tiling = tiling,
   .usage = VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT,
   // ...
};

vkCreateImage(device, &imageCreateInfo, nullptr, &fdmImage);

VkImageViewCreateInfo viewCreateInfo =
{
   .sType = VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO,
   .pNext = nullptr,
   .flags = 0,                      // VkImageViewCreateFlags
   .image = fdmImage,
   .viewType = VK_IMAGE_VIEW_TYPE_2D_ARRAY,
   .format = VK_FORMAT_R8G8_UNORM,
   .components = { 0 },             // VK_COMPONENT_SWIZZLE_IDENTITY
   .subresourceRange = {
        .aspectMask = VK_IMAGE_ASPECT_COLOR_BIT,
        .baseMipLevel = 0,
        .levelCount = 1,
        .baseArrayLayer = 0,
        .layerCount = 2,
   }
};

vkCreateImageView(device, &viewCreateInfo, nullptr, &fdmImageView);

// Add fdmImage to render pass

VkAttachmentReference fragmentDensityMapAttachmentReference =
{
   fdmAttachmentIdx,
   VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT,
};

VkRenderPassFragmentDensityMapCreateInfoEXT fdmAttachmentCreateInfo =
{
   VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT,
   // ...
   fragmentDensityMapAttachmentReference,
};

VkRenderPassCreateInfo2 renderPassCreateInfo =
{
   VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2,
   &fdmAttachmentCreateInfo,
   // ...
};

vkCreateRenderPass2(device, &renderPassCreateInfo, nullptr, &renderPass);

// Add fdmImage to framebuffer
// Color attachments can be created with VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT
// All attachments must be created with VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSETS_BIT_EXT
VkFramebufferCreateInfo framebufferCreateInfo =
{
   .sType = VK_STRUCTURE_TYPE_FRAME_BUFFER_CREATE_INFO,
   // ...
   .renderPass = renderPass,
   // ...
   .pAttachments = pAttachments, // Includes fdmImageView at fdmAttachmentIdx
   .width = 1024,
   .height = 1024,
   .layers = 1
};

vkCreateFramebuffer(device, &framebufferCreateInfo, nullptr, &framebuffer);

// Start recording render pass in command buffer

VkRenderPassBeginInfo renderPassBeginInfo =
{
   .sType = VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO,
   // ...
   .renderPass = renderPass,
   .framebuffer = framebuffer,
   // ...
};

// Can no longer modify the fdmImage's contents after this call
vkCmdBeginRenderPass2(commandBuffer, &renderPassBeginInfo, pSubpassBeginInfo);
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-09-25 (Matthew Netsch)

  - Initial version

- Revision 2, 2021-09-30 (Jon Leech)

  - Add interaction with
    [`VK_KHR_format_feature_flags2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)
    to `vk.xml`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_fragment_density_map"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
