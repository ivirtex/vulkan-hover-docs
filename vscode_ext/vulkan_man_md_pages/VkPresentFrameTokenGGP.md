# VkPresentFrameTokenGGP(3) Manual Page

## Name

VkPresentFrameTokenGGP - The Google Games Platform frame token



## <a href="#_c_specification" class="anchor"></a>C Specification

When the [`VK_GGP_frame_token`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GGP_frame_token.html) extension is
enabled, a Google Games Platform frame token **can** be specified when
presenting an image to a swapchain by adding a `VkPresentFrameTokenGGP`
structure to the `pNext` chain of the `VkPresentInfoKHR` structure.

The `VkPresentFrameTokenGGP` structure is defined as:

``` c
// Provided by VK_GGP_frame_token
typedef struct VkPresentFrameTokenGGP {
    VkStructureType    sType;
    const void*        pNext;
    GgpFrameToken      frameToken;
} VkPresentFrameTokenGGP;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `frameToken` is the Google Games Platform frame token.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPresentFrameTokenGGP-frameToken-02680"
  id="VUID-VkPresentFrameTokenGGP-frameToken-02680"></a>
  VUID-VkPresentFrameTokenGGP-frameToken-02680  
  `frameToken` **must** be a valid `GgpFrameToken`

Valid Usage (Implicit)

- <a href="#VUID-VkPresentFrameTokenGGP-sType-sType"
  id="VUID-VkPresentFrameTokenGGP-sType-sType"></a>
  VUID-VkPresentFrameTokenGGP-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_GGP_frame_token](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GGP_frame_token.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentFrameTokenGGP"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
