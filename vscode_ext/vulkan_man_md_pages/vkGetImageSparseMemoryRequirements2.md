# vkGetImageSparseMemoryRequirements2(3) Manual Page

## Name

vkGetImageSparseMemoryRequirements2 - Query the memory requirements for
a sparse image



## <a href="#_c_specification" class="anchor"></a>C Specification

To query sparse memory requirements for an image, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetImageSparseMemoryRequirements2(
    VkDevice                                    device,
    const VkImageSparseMemoryRequirementsInfo2* pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_memory_requirements2
void vkGetImageSparseMemoryRequirements2KHR(
    VkDevice                                    device,
    const VkImageSparseMemoryRequirementsInfo2* pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `pInfo` is a pointer to a `VkImageSparseMemoryRequirementsInfo2`
  structure containing parameters required for the memory requirements
  query.

- `pSparseMemoryRequirementCount` is a pointer to an integer related to
  the number of sparse memory requirements available or queried, as
  described below.

- `pSparseMemoryRequirements` is either `NULL` or a pointer to an array
  of `VkSparseImageMemoryRequirements2` structures.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageSparseMemoryRequirements2-device-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements2-device-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageSparseMemoryRequirements2-pInfo-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements2-pInfo-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements2-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSparseMemoryRequirementsInfo2.html)
  structure

- <a
  href="#VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirementCount-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirementCount-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirementCount-parameter  
  `pSparseMemoryRequirementCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirements-parameter"
  id="VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirements-parameter"></a>
  VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirements-parameter  
  If the value referenced by `pSparseMemoryRequirementCount` is not `0`,
  and `pSparseMemoryRequirements` is not `NULL`,
  `pSparseMemoryRequirements` **must** be a valid pointer to an array of
  `pSparseMemoryRequirementCount`
  [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSparseMemoryRequirementsInfo2.html),
[VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageSparseMemoryRequirements2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
