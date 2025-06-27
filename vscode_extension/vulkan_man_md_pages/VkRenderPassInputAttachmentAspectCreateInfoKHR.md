# VkRenderPassInputAttachmentAspectCreateInfo(3) Manual Page

## Name

VkRenderPassInputAttachmentAspectCreateInfo - Structure specifying, for
a given subpass/input attachment pair, which aspect can: be read.



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassInputAttachmentAspectCreateInfo` structure is defined
as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkRenderPassInputAttachmentAspectCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   aspectReferenceCount;
    const VkInputAttachmentAspectReference*    pAspectReferences;
} VkRenderPassInputAttachmentAspectCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance2
typedef VkRenderPassInputAttachmentAspectCreateInfo VkRenderPassInputAttachmentAspectCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `aspectReferenceCount` is the number of elements in the
  `pAspectReferences` array.

- `pAspectReferences` is a pointer to an array of `aspectReferenceCount`
  [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReference.html)
  structures containing a mask describing which aspect(s) **can** be
  accessed for a given input attachment within a given subpass.

## <a href="#_description" class="anchor"></a>Description

To specify which aspects of an input attachment **can** be read, add a
[VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)
structure to the `pNext` chain of the
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) structure:

An application **can** access any aspect of an input attachment that
does not have a specified aspect mask in the `pAspectReferences` array.
Otherwise, an application **must** not access aspect(s) of an input
attachment other than those in its specified aspect mask.

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassInputAttachmentAspectCreateInfo-sType-sType"
  id="VUID-VkRenderPassInputAttachmentAspectCreateInfo-sType-sType"></a>
  VUID-VkRenderPassInputAttachmentAspectCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO`

- <a
  href="#VUID-VkRenderPassInputAttachmentAspectCreateInfo-pAspectReferences-parameter"
  id="VUID-VkRenderPassInputAttachmentAspectCreateInfo-pAspectReferences-parameter"></a>
  VUID-VkRenderPassInputAttachmentAspectCreateInfo-pAspectReferences-parameter  
  `pAspectReferences` **must** be a valid pointer to an array of
  `aspectReferenceCount` valid
  [VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReference.html)
  structures

- <a
  href="#VUID-VkRenderPassInputAttachmentAspectCreateInfo-aspectReferenceCount-arraylength"
  id="VUID-VkRenderPassInputAttachmentAspectCreateInfo-aspectReferenceCount-arraylength"></a>
  VUID-VkRenderPassInputAttachmentAspectCreateInfo-aspectReferenceCount-arraylength  
  `aspectReferenceCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkInputAttachmentAspectReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReference.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassInputAttachmentAspectCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
