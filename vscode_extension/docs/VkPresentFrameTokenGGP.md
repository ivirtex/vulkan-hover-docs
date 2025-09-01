# VkPresentFrameTokenGGP(3) Manual Page

## Name

VkPresentFrameTokenGGP - The Google Games Platform frame token



## [](#_c_specification)C Specification

When the `VK_GGP_frame_token` extension is enabled, a Google Games Platform frame token **can** be specified when presenting an image to a swapchain by adding a `VkPresentFrameTokenGGP` structure to the `pNext` chain of the `VkPresentInfoKHR` structure.

The `VkPresentFrameTokenGGP` structure is defined as:

```c++
// Provided by VK_GGP_frame_token
typedef struct VkPresentFrameTokenGGP {
    VkStructureType    sType;
    const void*        pNext;
    GgpFrameToken      frameToken;
} VkPresentFrameTokenGGP;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `frameToken` is the Google Games Platform frame token.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPresentFrameTokenGGP-frameToken-02680)VUID-VkPresentFrameTokenGGP-frameToken-02680  
  `frameToken` **must** be a valid `GgpFrameToken`

Valid Usage (Implicit)

- [](#VUID-VkPresentFrameTokenGGP-sType-sType)VUID-VkPresentFrameTokenGGP-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP`

## [](#_see_also)See Also

[VK\_GGP\_frame\_token](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_GGP_frame_token.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentFrameTokenGGP).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0