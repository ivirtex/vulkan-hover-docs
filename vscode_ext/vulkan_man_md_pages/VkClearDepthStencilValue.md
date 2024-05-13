# VkClearDepthStencilValue(3) Manual Page

## Name

VkClearDepthStencilValue - Structure specifying a clear depth stencil
value



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkClearDepthStencilValue` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkClearDepthStencilValue {
    float       depth;
    uint32_t    stencil;
} VkClearDepthStencilValue;
```

## <a href="#_members" class="anchor"></a>Members

- `depth` is the clear value for the depth aspect of the depth/stencil
  attachment. It is a floating-point value which is automatically
  converted to the attachment’s format.

- `stencil` is the clear value for the stencil aspect of the
  depth/stencil attachment. It is a 32-bit integer value which is
  converted to the attachment’s format by taking the appropriate number
  of LSBs.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkClearDepthStencilValue-depth-00022"
  id="VUID-VkClearDepthStencilValue-depth-00022"></a>
  VUID-VkClearDepthStencilValue-depth-00022  
  Unless the
  [`VK_EXT_depth_range_unrestricted`](VK_EXT_depth_range_unrestricted.html)
  extension is enabled `depth` **must** be between `0.0` and `1.0`,
  inclusive

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html),
[vkCmdClearDepthStencilImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdClearDepthStencilImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkClearDepthStencilValue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
