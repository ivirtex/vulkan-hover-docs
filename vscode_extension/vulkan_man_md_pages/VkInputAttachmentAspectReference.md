# VkInputAttachmentAspectReference(3) Manual Page

## Name

VkInputAttachmentAspectReference - Structure specifying a subpass/input
attachment pair and an aspect mask that can: be read.



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkInputAttachmentAspectReference` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkInputAttachmentAspectReference {
    uint32_t              subpass;
    uint32_t              inputAttachmentIndex;
    VkImageAspectFlags    aspectMask;
} VkInputAttachmentAspectReference;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance2
typedef VkInputAttachmentAspectReference VkInputAttachmentAspectReferenceKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `subpass` is an index into the `pSubpasses` array of the parent
  `VkRenderPassCreateInfo` structure.

- `inputAttachmentIndex` is an index into the `pInputAttachments` of the
  specified subpass.

- `aspectMask` is a mask of which aspect(s) **can** be accessed within
  the specified subpass.

## <a href="#_description" class="anchor"></a>Description

This structure specifies an aspect mask for a specific input attachment
of a specific subpass in the render pass.

`subpass` and `inputAttachmentIndex` index into the render pass as:

``` c
pCreateInfo->pSubpasses[subpass].pInputAttachments[inputAttachmentIndex]
```

Valid Usage

- <a href="#VUID-VkInputAttachmentAspectReference-aspectMask-01964"
  id="VUID-VkInputAttachmentAspectReference-aspectMask-01964"></a>
  VUID-VkInputAttachmentAspectReference-aspectMask-01964  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_METADATA_BIT`

- <a href="#VUID-VkInputAttachmentAspectReference-aspectMask-02250"
  id="VUID-VkInputAttachmentAspectReference-aspectMask-02250"></a>
  VUID-VkInputAttachmentAspectReference-aspectMask-02250  
  `aspectMask` **must** not include
  `VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` for any index *i*

Valid Usage (Implicit)

- <a href="#VUID-VkInputAttachmentAspectReference-aspectMask-parameter"
  id="VUID-VkInputAttachmentAspectReference-aspectMask-parameter"></a>
  VUID-VkInputAttachmentAspectReference-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) values

- <a
  href="#VUID-VkInputAttachmentAspectReference-aspectMask-requiredbitmask"
  id="VUID-VkInputAttachmentAspectReference-aspectMask-requiredbitmask"></a>
  VUID-VkInputAttachmentAspectReference-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkInputAttachmentAspectReference"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
