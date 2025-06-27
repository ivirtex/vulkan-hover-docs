# VkClearAttachment(3) Manual Page

## Name

VkClearAttachment - Structure specifying a clear attachment



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkClearAttachment` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkClearAttachment {
    VkImageAspectFlags    aspectMask;
    uint32_t              colorAttachment;
    VkClearValue          clearValue;
} VkClearAttachment;
```

## <a href="#_members" class="anchor"></a>Members

- `aspectMask` is a mask selecting the color, depth and/or stencil
  aspects of the attachment to be cleared.

- `colorAttachment` is only meaningful if `VK_IMAGE_ASPECT_COLOR_BIT` is
  set in `aspectMask`, in which case it is an index into the currently
  bound color attachments.

- `clearValue` is the color or depth/stencil value to clear the
  attachment to, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#clears-values"
  target="_blank" rel="noopener">Clear Values</a> below.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkClearAttachment-aspectMask-00019"
  id="VUID-VkClearAttachment-aspectMask-00019"></a>
  VUID-VkClearAttachment-aspectMask-00019  
  If `aspectMask` includes `VK_IMAGE_ASPECT_COLOR_BIT`, it **must** not
  include `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkClearAttachment-aspectMask-00020"
  id="VUID-VkClearAttachment-aspectMask-00020"></a>
  VUID-VkClearAttachment-aspectMask-00020  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_METADATA_BIT`

- <a href="#VUID-VkClearAttachment-aspectMask-02246"
  id="VUID-VkClearAttachment-aspectMask-02246"></a>
  VUID-VkClearAttachment-aspectMask-02246  
  `aspectMask` **must** not include
  `VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` for any index *i*

Valid Usage (Implicit)

- <a href="#VUID-VkClearAttachment-aspectMask-parameter"
  id="VUID-VkClearAttachment-aspectMask-parameter"></a>
  VUID-VkClearAttachment-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) values

- <a href="#VUID-VkClearAttachment-aspectMask-requiredbitmask"
  id="VUID-VkClearAttachment-aspectMask-requiredbitmask"></a>
  VUID-VkClearAttachment-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearAttachments.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkClearAttachment"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
