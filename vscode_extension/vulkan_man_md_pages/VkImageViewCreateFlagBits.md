# VkImageViewCreateFlagBits(3) Manual Page

## Name

VkImageViewCreateFlagBits - Bitmask specifying additional parameters of
an image view



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`flags`, specifying
additional parameters of an image view, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkImageViewCreateFlagBits {
  // Provided by VK_EXT_fragment_density_map
    VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT = 0x00000001,
  // Provided by VK_EXT_descriptor_buffer
    VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT = 0x00000004,
  // Provided by VK_EXT_fragment_density_map2
    VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT = 0x00000002,
} VkImageViewCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` specifies
  that the fragment density map will be read by device during
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT` specifies
  that the fragment density map will be read by the host during
  [vkEndCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEndCommandBuffer.html) for the primary command
  buffer that the render pass is recorded into

- `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
  specifies that the image view **can** be used with descriptor buffers
  when capturing and replaying (e.g. for trace capture and replay), see
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  for more detail.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkImageViewCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
