# VkClearAttachment(3) Manual Page

## Name

VkClearAttachment - Structure specifying a clear attachment



## [](#_c_specification)C Specification

The `VkClearAttachment` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkClearAttachment {
    VkImageAspectFlags    aspectMask;
    uint32_t              colorAttachment;
    VkClearValue          clearValue;
} VkClearAttachment;
```

## [](#_members)Members

- `aspectMask` is a mask selecting the color, depth and/or stencil aspects of the attachment to be cleared.
- `colorAttachment` is only meaningful if `VK_IMAGE_ASPECT_COLOR_BIT` is set in `aspectMask`, in which case it is an index into the bound color attachments.
- `clearValue` is the color or depth/stencil value to clear the attachment to, as described in [Clear Values](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#clears-values) below.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClearAttachment-aspectMask-00019)VUID-VkClearAttachment-aspectMask-00019  
  If `aspectMask` includes `VK_IMAGE_ASPECT_COLOR_BIT`, it **must** not include `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkClearAttachment-aspectMask-00020)VUID-VkClearAttachment-aspectMask-00020  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_METADATA_BIT`
- [](#VUID-VkClearAttachment-aspectMask-02246)VUID-VkClearAttachment-aspectMask-02246  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT` for any index *i*

Valid Usage (Implicit)

- [](#VUID-VkClearAttachment-aspectMask-parameter)VUID-VkClearAttachment-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) values
- [](#VUID-VkClearAttachment-aspectMask-requiredbitmask)VUID-VkClearAttachment-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [vkCmdClearAttachments](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearAttachments.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClearAttachment).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0