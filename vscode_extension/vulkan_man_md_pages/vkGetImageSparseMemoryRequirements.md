# vkGetImageSparseMemoryRequirements(3) Manual Page

## Name

vkGetImageSparseMemoryRequirements - Query the memory requirements for a
sparse image



## <a href="#_c_specification" class="anchor"></a>C Specification

To query sparse memory requirements for an image, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetImageSparseMemoryRequirements(
    VkDevice                                    device,
    VkImage                                     image,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements*            pSparseMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `image` is the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object to get the memory
  requirements for.

- `pSparseMemoryRequirementCount` is a pointer to an integer related to
  the number of sparse memory requirements available or queried, as
  described below.

- `pSparseMemoryRequirements` is either `NULL` or a pointer to an array
  of `VkSparseImageMemoryRequirements` structures.

## <a href="#_description" class="anchor"></a>Description

If `pSparseMemoryRequirements` is `NULL`, then the number of sparse
memory requirements available is returned in
`pSparseMemoryRequirementCount`. Otherwise,
`pSparseMemoryRequirementCount` **must** point to a variable set by the
application to the number of elements in the `pSparseMemoryRequirements`
array, and on return the variable is overwritten with the number of
structures actually written to `pSparseMemoryRequirements`. If
`pSparseMemoryRequirementCount` is less than the number of sparse memory
requirements available, at most `pSparseMemoryRequirementCount`
structures will be written.

If the image was not created with `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`
then `pSparseMemoryRequirementCount` will be set to zero and
`pSparseMemoryRequirements` will not be written to.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>It is legal for an implementation to report a larger value in
<code>VkMemoryRequirements</code>::<code>size</code> than would be
obtained by adding together memory sizes for all
<code>VkSparseImageMemoryRequirements</code> returned by
<code>vkGetImageSparseMemoryRequirements</code>. This
<strong>may</strong> occur when the implementation requires unused
padding in the address range describing the resource.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageSparseMemoryRequirements-device-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements-device-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageSparseMemoryRequirements-image-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements-image-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a
  href="#VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter  
  `pSparseMemoryRequirementCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter  
  If the value referenced by `pSparseMemoryRequirementCount` is not `0`,
  and `pSparseMemoryRequirements` is not `NULL`,
  `pSparseMemoryRequirements` **must** be a valid pointer to an array of
  `pSparseMemoryRequirementCount`
  [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements.html)
  structures

- <a href="#VUID-vkGetImageSparseMemoryRequirements-image-parent"
  id="VUID-vkGetImageSparseMemoryRequirements-image-parent"></a>
  VUID-vkGetImageSparseMemoryRequirements-image-parent  
  `image` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageSparseMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
