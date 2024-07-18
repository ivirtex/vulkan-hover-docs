# vkGetBufferCollectionPropertiesFUCHSIA(3) Manual Page

## Name

vkGetBufferCollectionPropertiesFUCHSIA - Retrieve properties from a
buffer collection



## <a href="#_c_specification" class="anchor"></a>C Specification

After constraints have been set on the buffer collection by calling
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)
or
[vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html),
call `vkGetBufferCollectionPropertiesFUCHSIA` to retrieve the negotiated
and finalized properties of the buffer collection.

The call to `vkGetBufferCollectionPropertiesFUCHSIA` is synchronous. It
waits for the Sysmem format negotiation and buffer collection allocation
to complete before returning.

``` c
// Provided by VK_FUCHSIA_buffer_collection
VkResult vkGetBufferCollectionPropertiesFUCHSIA(
    VkDevice                                    device,
    VkBufferCollectionFUCHSIA                   collection,
    VkBufferCollectionPropertiesFUCHSIA*        pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device handle

- `collection` is the
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- `pProperties` is a pointer to the retrieved
  [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)
  struct

## <a href="#_description" class="anchor"></a>Description

For image-based buffer collections, upon calling
`vkGetBufferCollectionPropertiesFUCHSIA`, Sysmem will choose an element
of the
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)::`pImageCreateInfos`
established by the preceding call to
[vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html).
The index of the element chosen is stored in and can be retrieved from
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`createInfoIndex`.

For buffer-based buffer collections, a single
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) is specified as
[VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferConstraintsInfoFUCHSIA.html)::`createInfo`.
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)::`createInfoIndex`
will therefore always be zero.

`vkGetBufferCollectionPropertiesFUCHSIA` **may** fail if Sysmem is
unable to resolve the constraints of all of the participants in the
buffer collection. If that occurs,
`vkGetBufferCollectionPropertiesFUCHSIA` will return
`VK_ERROR_INITIALIZATION_FAILED`.

Valid Usage

- <a href="#VUID-vkGetBufferCollectionPropertiesFUCHSIA-None-06405"
  id="VUID-vkGetBufferCollectionPropertiesFUCHSIA-None-06405"></a>
  VUID-vkGetBufferCollectionPropertiesFUCHSIA-None-06405  
  Prior to calling
  [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html),
  the constraints on the buffer collection **must** have been set by
  either
  [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)
  or
  [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

Valid Usage (Implicit)

- <a href="#VUID-vkGetBufferCollectionPropertiesFUCHSIA-device-parameter"
  id="VUID-vkGetBufferCollectionPropertiesFUCHSIA-device-parameter"></a>
  VUID-vkGetBufferCollectionPropertiesFUCHSIA-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parameter"
  id="VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parameter"></a>
  VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parameter  
  `collection` **must** be a valid
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) handle

- <a
  href="#VUID-vkGetBufferCollectionPropertiesFUCHSIA-pProperties-parameter"
  id="VUID-vkGetBufferCollectionPropertiesFUCHSIA-pProperties-parameter"></a>
  VUID-vkGetBufferCollectionPropertiesFUCHSIA-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)
  structure

- <a href="#VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parent"
  id="VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parent"></a>
  VUID-vkGetBufferCollectionPropertiesFUCHSIA-collection-parent  
  `collection` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetBufferCollectionPropertiesFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
