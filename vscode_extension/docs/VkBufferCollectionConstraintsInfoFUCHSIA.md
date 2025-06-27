# VkBufferCollectionConstraintsInfoFUCHSIA(3) Manual Page

## Name

VkBufferCollectionConstraintsInfoFUCHSIA - Structure of general buffer
collection constraints



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCollectionConstraintsInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferCollectionConstraintsInfoFUCHSIA {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           minBufferCount;
    uint32_t           maxBufferCount;
    uint32_t           minBufferCountForCamping;
    uint32_t           minBufferCountForDedicatedSlack;
    uint32_t           minBufferCountForSharedSlack;
} VkBufferCollectionConstraintsInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `minBufferCount` is the minimum number of buffers available in the
  collection

- `maxBufferCount` is the maximum number of buffers allowed in the
  collection

- `minBufferCountForCamping` is the per-participant minimum buffers for
  camping

- `minBufferCountForDedicatedSlack` is the per-participant minimum
  buffers for dedicated slack

- `minBufferCountForSharedSlack` is the per-participant minimum buffers
  for shared slack

## <a href="#_description" class="anchor"></a>Description

Sysmem uses all buffer count parameters in combination to determine the
number of buffers it will allocate. Sysmem defines buffer count
constraints in `fuchsia.sysmem/constraints.fidl`.

*Camping* as referred to by `minBufferCountForCamping`, is the number of
buffers that should be available for the participant that are not for
transient use. This number of buffers is required for the participant to
logically operate.

*Slack* as referred to by `minBufferCountForDedicatedSlack` and
`minBufferCountForSharedSlack`, refers to the number of buffers desired
by participants for optimal performance.
`minBufferCountForDedicatedSlack` refers to the current participant.
`minBufferCountForSharedSlack` refers to buffer slack for all
participants in the collection.

Valid Usage (Implicit)

- <a href="#VUID-VkBufferCollectionConstraintsInfoFUCHSIA-sType-sType"
  id="VUID-VkBufferCollectionConstraintsInfoFUCHSIA-sType-sType"></a>
  VUID-VkBufferCollectionConstraintsInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_CONSTRAINTS_INFO_FUCHSIA`

- <a href="#VUID-VkBufferCollectionConstraintsInfoFUCHSIA-pNext-pNext"
  id="VUID-VkBufferCollectionConstraintsInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkBufferCollectionConstraintsInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferConstraintsInfoFUCHSIA.html),
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCollectionConstraintsInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
