# vkGetFramebufferTilePropertiesQCOM(3) Manual Page

## Name

vkGetFramebufferTilePropertiesQCOM - Get tile properties from the
attachments in framebuffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the tile properties from the attachments in framebuffer, call:

``` c
// Provided by VK_QCOM_tile_properties
VkResult vkGetFramebufferTilePropertiesQCOM(
    VkDevice                                    device,
    VkFramebuffer                               framebuffer,
    uint32_t*                                   pPropertiesCount,
    VkTilePropertiesQCOM*                       pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is a logical device associated with the framebuffer.

- `framebuffer` is a handle of the framebuffer to query.

- `pPropertiesCount` is a pointer to an integer related to the number of
  tile properties available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html) structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of tile properties available
is returned in `pPropertiesCount`. Otherwise, `pPropertiesCount`
**must** point to a variable set by the user to the number of elements
in the `pProperties` array, and on return the variable is overwritten
with the number of properties actually written to `pProperties`. If
`pPropertiesCount` is less than the number of tile properties available,
at most `pPropertiesCount` structures will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available properties were returned.

The number of tile properties available is determined by the number of
merged subpasses, and each tile property is associated with a merged
subpass. There will be at most as many properties as there are subpasses
within the render pass. To obtain the tile properties for a given merged
subpass, the `pProperties` array can be indexed using the
`postMergeIndex` value provided in
[VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackInfoEXT.html).

Valid Usage (Implicit)

- <a href="#VUID-vkGetFramebufferTilePropertiesQCOM-device-parameter"
  id="VUID-vkGetFramebufferTilePropertiesQCOM-device-parameter"></a>
  VUID-vkGetFramebufferTilePropertiesQCOM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetFramebufferTilePropertiesQCOM-framebuffer-parameter"
  id="VUID-vkGetFramebufferTilePropertiesQCOM-framebuffer-parameter"></a>
  VUID-vkGetFramebufferTilePropertiesQCOM-framebuffer-parameter  
  `framebuffer` **must** be a valid [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)
  handle

- <a
  href="#VUID-vkGetFramebufferTilePropertiesQCOM-pPropertiesCount-parameter"
  id="VUID-vkGetFramebufferTilePropertiesQCOM-pPropertiesCount-parameter"></a>
  VUID-vkGetFramebufferTilePropertiesQCOM-pPropertiesCount-parameter  
  `pPropertiesCount` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkGetFramebufferTilePropertiesQCOM-pProperties-parameter"
  id="VUID-vkGetFramebufferTilePropertiesQCOM-pProperties-parameter"></a>
  VUID-vkGetFramebufferTilePropertiesQCOM-pProperties-parameter  
  If the value referenced by `pPropertiesCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertiesCount`
  [VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html) structures

- <a href="#VUID-vkGetFramebufferTilePropertiesQCOM-framebuffer-parent"
  id="VUID-vkGetFramebufferTilePropertiesQCOM-framebuffer-parent"></a>
  VUID-vkGetFramebufferTilePropertiesQCOM-framebuffer-parent  
  `framebuffer` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

- `VK_INCOMPLETE`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_tile_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_tile_properties.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html),
[VkTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTilePropertiesQCOM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetFramebufferTilePropertiesQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
