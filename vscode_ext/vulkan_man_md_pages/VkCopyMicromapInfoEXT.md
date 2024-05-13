# VkCopyMicromapInfoEXT(3) Manual Page

## Name

VkCopyMicromapInfoEXT - Parameters for copying a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyMicromapInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkCopyMicromapInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkMicromapEXT            src;
    VkMicromapEXT            dst;
    VkCopyMicromapModeEXT    mode;
} VkCopyMicromapInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `src` is the source micromap for the copy.

- `dst` is the target micromap for the copy.

- `mode` is a [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html) value
  specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyMicromapInfoEXT-mode-07531"
  id="VUID-VkCopyMicromapInfoEXT-mode-07531"></a>
  VUID-VkCopyMicromapInfoEXT-mode-07531  
  `mode` **must** be `VK_COPY_MICROMAP_MODE_COMPACT_EXT` or
  `VK_COPY_MICROMAP_MODE_CLONE_EXT`

- <a href="#VUID-VkCopyMicromapInfoEXT-src-07532"
  id="VUID-VkCopyMicromapInfoEXT-src-07532"></a>
  VUID-VkCopyMicromapInfoEXT-src-07532  
  The source acceleration structure `src` **must** have been constructed
  prior to the execution of this command

- <a href="#VUID-VkCopyMicromapInfoEXT-mode-07533"
  id="VUID-VkCopyMicromapInfoEXT-mode-07533"></a>
  VUID-VkCopyMicromapInfoEXT-mode-07533  
  If `mode` is `VK_COPY_MICROMAP_MODE_COMPACT_EXT`, `src` **must** have
  been constructed with `VK_BUILD_MICROMAP_ALLOW_COMPACTION_BIT_EXT` in
  the build

- <a href="#VUID-VkCopyMicromapInfoEXT-buffer-07534"
  id="VUID-VkCopyMicromapInfoEXT-buffer-07534"></a>
  VUID-VkCopyMicromapInfoEXT-buffer-07534  
  The `buffer` used to create `src` **must** be bound to device memory

- <a href="#VUID-VkCopyMicromapInfoEXT-buffer-07535"
  id="VUID-VkCopyMicromapInfoEXT-buffer-07535"></a>
  VUID-VkCopyMicromapInfoEXT-buffer-07535  
  The `buffer` used to create `dst` **must** be bound to device memory

Valid Usage (Implicit)

- <a href="#VUID-VkCopyMicromapInfoEXT-sType-sType"
  id="VUID-VkCopyMicromapInfoEXT-sType-sType"></a>
  VUID-VkCopyMicromapInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MICROMAP_INFO_EXT`

- <a href="#VUID-VkCopyMicromapInfoEXT-pNext-pNext"
  id="VUID-VkCopyMicromapInfoEXT-pNext-pNext"></a>
  VUID-VkCopyMicromapInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyMicromapInfoEXT-src-parameter"
  id="VUID-VkCopyMicromapInfoEXT-src-parameter"></a>
  VUID-VkCopyMicromapInfoEXT-src-parameter  
  `src` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

- <a href="#VUID-VkCopyMicromapInfoEXT-dst-parameter"
  id="VUID-VkCopyMicromapInfoEXT-dst-parameter"></a>
  VUID-VkCopyMicromapInfoEXT-dst-parameter  
  `dst` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

- <a href="#VUID-VkCopyMicromapInfoEXT-mode-parameter"
  id="VUID-VkCopyMicromapInfoEXT-mode-parameter"></a>
  VUID-VkCopyMicromapInfoEXT-mode-parameter  
  `mode` **must** be a valid
  [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html) value

- <a href="#VUID-VkCopyMicromapInfoEXT-commonparent"
  id="VUID-VkCopyMicromapInfoEXT-commonparent"></a>
  VUID-VkCopyMicromapInfoEXT-commonparent  
  Both of `dst`, and `src` **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html),
[VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapEXT.html),
[vkCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMicromapInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
