# VkRenderPassFragmentDensityMapCreateInfoEXT(3) Manual Page

## Name

VkRenderPassFragmentDensityMapCreateInfoEXT - Structure containing
fragment density map attachment for render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

If the [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)::`pNext`
chain includes a `VkRenderPassFragmentDensityMapCreateInfoEXT`
structure, then that structure includes a fragment density map
attachment for the render pass.

The `VkRenderPassFragmentDensityMapCreateInfoEXT` structure is defined
as:

``` c
// Provided by VK_EXT_fragment_density_map
typedef struct VkRenderPassFragmentDensityMapCreateInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkAttachmentReference    fragmentDensityMapAttachment;
} VkRenderPassFragmentDensityMapCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fragmentDensityMapAttachment` is the fragment density map to use for
  the render pass.

## <a href="#_description" class="anchor"></a>Description

The fragment density map is read at an implementation-dependent time
with the following constraints determined by the attachment’s image view
`flags`:

- `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` specifies
  that the fragment density map will be read by the device during
  `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DEFERRED_BIT_EXT` specifies
  that the fragment density map will be read by the host during
  [vkEndCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEndCommandBuffer.html) of the primary command
  buffer that the render pass is recorded into

- Otherwise the fragment density map will be read by the host during
  [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html)

The fragment density map **may** additionally be read by the device
during `VK_PIPELINE_STAGE_FRAGMENT_DENSITY_PROCESS_BIT_EXT` for any
mode.

If this structure is not present, it is as if
`fragmentDensityMapAttachment` was given as `VK_ATTACHMENT_UNUSED`.

Valid Usage

- <a
  href="#VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02548"
  id="VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02548"></a>
  VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02548  
  If `fragmentDensityMapAttachment` is not `VK_ATTACHMENT_UNUSED`,
  `fragmentDensityMapAttachment` **must** not be an element of
  `VkSubpassDescription`::`pInputAttachments`,
  `VkSubpassDescription`::`pColorAttachments`,
  `VkSubpassDescription`::`pResolveAttachments`,
  `VkSubpassDescription`::`pDepthStencilAttachment`, or
  `VkSubpassDescription`::`pPreserveAttachments` for any subpass

- <a
  href="#VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02549"
  id="VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02549"></a>
  VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02549  
  If `fragmentDensityMapAttachment` is not `VK_ATTACHMENT_UNUSED`,
  `layout` **must** be equal to
  `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`, or
  `VK_IMAGE_LAYOUT_GENERAL`

- <a
  href="#VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02550"
  id="VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02550"></a>
  VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02550  
  If `fragmentDensityMapAttachment` is not `VK_ATTACHMENT_UNUSED`,
  `fragmentDensityMapAttachment` **must** reference an attachment with a
  `loadOp` equal to `VK_ATTACHMENT_LOAD_OP_LOAD` or
  `VK_ATTACHMENT_LOAD_OP_DONT_CARE`

- <a
  href="#VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02551"
  id="VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02551"></a>
  VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-02551  
  If `fragmentDensityMapAttachment` is not `VK_ATTACHMENT_UNUSED`,
  `fragmentDensityMapAttachment` **must** reference an attachment with a
  `storeOp` equal to `VK_ATTACHMENT_STORE_OP_DONT_CARE`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-sType-sType"
  id="VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-sType-sType"></a>
  VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_CREATE_INFO_EXT`

- <a
  href="#VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-parameter"
  id="VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-parameter"></a>
  VUID-VkRenderPassFragmentDensityMapCreateInfoEXT-fragmentDensityMapAttachment-parameter  
  `fragmentDensityMapAttachment` **must** be a valid
  [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html),
[VkAttachmentReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassFragmentDensityMapCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
