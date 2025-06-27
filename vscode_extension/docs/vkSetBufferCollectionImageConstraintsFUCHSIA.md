# vkSetBufferCollectionImageConstraintsFUCHSIA(3) Manual Page

## Name

vkSetBufferCollectionImageConstraintsFUCHSIA - Set image-based
constraints for a buffer collection



## <a href="#_c_specification" class="anchor"></a>C Specification

Setting the constraints on the buffer collection initiates the format
negotiation and allocation of the buffer collection. To set the
constraints on a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) buffer collection, call:

``` c
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkSetBufferCollectionImageConstraintsFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    const VkImageConstraintsInfoFUCHSIA*        pImageConstraintsInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device

- `collection` is the
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- `pImageConstraintsInfo` is a pointer to a
  [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)
  structure

## <a href="#_description" class="anchor"></a>Description

`vkSetBufferCollectionImageConstraintsFUCHSIA` **may** fail if
`pImageConstraintsInfo->formatConstraintsCount` is larger than the
implementation-defined limit. If that occurs,
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)
will return `VK_ERROR_INITIALIZATION_FAILED`.

`vkSetBufferCollectionImageConstraintsFUCHSIA` **may** fail if the
implementation does not support any of the formats described by the
`pImageConstraintsInfo` structure. If that occurs,
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)
will return `VK_ERROR_FORMAT_NOT_SUPPORTED`.

Valid Usage

- <a
  href="#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-06394"
  id="VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-06394"></a>
  VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-06394  
  `vkSetBufferCollectionImageConstraintsFUCHSIA` or
  `vkSetBufferCollectionBufferConstraintsFUCHSIA` **must** not have
  already been called on `collection`

Valid Usage (Implicit)

- <a
  href="#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-device-parameter"
  id="VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-device-parameter"></a>
  VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parameter"
  id="VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parameter"></a>
  VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parameter  
  `collection` **must** be a valid
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- <a
  href="#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-pImageConstraintsInfo-parameter"
  id="VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-pImageConstraintsInfo-parameter"></a>
  VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-pImageConstraintsInfo-parameter  
  `pImageConstraintsInfo` **must** be a valid pointer to a valid
  [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)
  structure

- <a
  href="#VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parent"
  id="VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parent"></a>
  VUID-vkSetBufferCollectionImageConstraintsFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_FORMAT_NOT_SUPPORTED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetBufferCollectionImageConstraintsFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
