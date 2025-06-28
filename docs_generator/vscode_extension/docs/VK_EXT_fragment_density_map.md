# VK\_EXT\_fragment\_density\_map(3) Manual Page

## Name

VK\_EXT\_fragment\_density\_map - device extension



## [](#_registered_extension_number)Registered Extension Number

219

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_dynamic\_rendering
- Interacts with VK\_KHR\_format\_feature\_flags2

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_fragment\_invocation\_density](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_fragment_invocation_density.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_fragment_density_map%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_fragment_density_map%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-09-30

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_fragment_invocation_density`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_fragment_invocation_density.txt)

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

## [](#_description)Description

This extension allows an application to specify areas of the render target where the fragment shader may be invoked fewer times. These fragments are broadcasted out to multiple pixels to cover the render target.

The primary use of this extension is to reduce workloads in areas where lower quality may not be perceived such as the distorted edges of a lens or the periphery of a user’s gaze.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentDensityMapFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapPropertiesEXT.html)
- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html):
  
  - [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html):
  
  - [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_FRAGMENT_DENSITY_MAP_EXTENSION_NAME`
- `VK_EXT_FRAGMENT_DENSITY_MAP_SPEC_VERSION`
- Extending [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html):
  
  - `VK_ACCESS_FRAGMENT_DENSITY_MAP_READ_BIT_EXT`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`
- Extending [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateFlagBits.html):
  
  - `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT`
- Extending [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html):
  
  - `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlagBits.html):
  
  - `VK_SAMPLER_CREATE_SUBSAMPLED_BIT_EXT`
  - `VK_SAMPLER_CREATE_SUBSAMPLED_COARSE_RECONSTRUCTION_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT`

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_FRAGMENT_DENSITY_MAP_BIT_EXT`

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html) is supported:

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
  - `VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_INFO_EXT`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`FragInvocationCountEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-fraginvocationcount)
- [`FragSizeEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-fragsize)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`FragmentDensityEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-FragmentDensityEXT)

## [](#_examples)Examples

### [](#_fragment_density_map)Fragment Density Map

An image can be bound as a fragment density map attachment to a render pass. This image contains normalized (x, y) float component fragment density values for regions of the framebuffer that will be used in rasterization for every subpass. A float component ranges from (0.0, 1.0] where 1.0 means full density along that axis. Implementations [use these values as hints](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragmentdensitymapops) to optimize rendering in areas of low density. Subpass color and depth attachments can be created as subsampled, which can help to further optimize rendering in areas of low density.

The density map image can be modified by the application until calling `vkCmdBeginRenderPass` for the render pass that uses the image. If `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` is used, then the application can modify the image until the device reads it during `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`.

```c++
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

## [](#_version_history)Version History

- Revision 1, 2018-09-25 (Matthew Netsch)
  
  - Initial version
- Revision 2, 2021-09-30 (Jon Leech)
  
  - Add interaction with `VK_KHR_format_feature_flags2` to `vk.xml`

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_fragment_density_map)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0