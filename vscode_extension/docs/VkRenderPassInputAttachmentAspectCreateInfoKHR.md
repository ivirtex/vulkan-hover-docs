# VkRenderPassInputAttachmentAspectCreateInfo(3) Manual Page

## Name

VkRenderPassInputAttachmentAspectCreateInfo - Structure specifying, for a given subpass/input attachment pair, which aspect can: be read.



## [](#_c_specification)C Specification

The `VkRenderPassInputAttachmentAspectCreateInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_1
typedef struct VkRenderPassInputAttachmentAspectCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   aspectReferenceCount;
    const VkInputAttachmentAspectReference*    pAspectReferences;
} VkRenderPassInputAttachmentAspectCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance2
typedef VkRenderPassInputAttachmentAspectCreateInfo VkRenderPassInputAttachmentAspectCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `aspectReferenceCount` is the number of elements in the `pAspectReferences` array.
- `pAspectReferences` is a pointer to an array of `aspectReferenceCount` [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInputAttachmentAspectReference.html) structures containing a mask describing which aspect(s) **can** be accessed for a given input attachment within a given subpass.

## [](#_description)Description

To specify which aspects of an input attachment **can** be read, add a [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html) structure to the `pNext` chain of the [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) structure:

An application **can** access any aspect of an input attachment that does not have a specified aspect mask in the `pAspectReferences` array. Otherwise, an application **must** not access aspect(s) of an input attachment other than those in its specified aspect mask.

Valid Usage (Implicit)

- [](#VUID-VkRenderPassInputAttachmentAspectCreateInfo-sType-sType)VUID-VkRenderPassInputAttachmentAspectCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO`
- [](#VUID-VkRenderPassInputAttachmentAspectCreateInfo-pAspectReferences-parameter)VUID-VkRenderPassInputAttachmentAspectCreateInfo-pAspectReferences-parameter  
  `pAspectReferences` **must** be a valid pointer to an array of `aspectReferenceCount` valid [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInputAttachmentAspectReference.html) structures
- [](#VUID-VkRenderPassInputAttachmentAspectCreateInfo-aspectReferenceCount-arraylength)VUID-VkRenderPassInputAttachmentAspectCreateInfo-aspectReferenceCount-arraylength  
  `aspectReferenceCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInputAttachmentAspectReference.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassInputAttachmentAspectCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0