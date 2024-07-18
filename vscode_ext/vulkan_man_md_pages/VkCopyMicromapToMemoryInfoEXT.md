# VkCopyMicromapToMemoryInfoEXT(3) Manual Page

## Name

VkCopyMicromapToMemoryInfoEXT - Parameters for serializing a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkCopyMicromapToMemoryInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkMicromapEXT               src;
    VkDeviceOrHostAddressKHR    dst;
    VkCopyMicromapModeEXT       mode;
} VkCopyMicromapToMemoryInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `src` is the source micromap for the copy

- `dst` is the device or host address to memory which is the target for
  the copy

- `mode` is a [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html) value
  specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-src-07540"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-src-07540"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-src-07540  
  The source micromap `src` **must** have been constructed prior to the
  execution of this command

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-dst-07541"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-dst-07541"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-dst-07541  
  The memory pointed to by `dst` **must** be at least as large as the
  serialization size of `src`, as reported by
  [vkWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWriteMicromapsPropertiesEXT.html) or
  [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteMicromapsPropertiesEXT.html)
  with a query type of `VK_QUERY_TYPE_MICROMAP_SERIALIZATION_SIZE_EXT`

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-mode-07542"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-mode-07542"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-mode-07542  
  `mode` **must** be `VK_COPY_MICROMAP_MODE_SERIALIZE_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-sType-sType"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-sType-sType"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COPY_MICROMAP_TO_MEMORY_INFO_EXT`

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-pNext-pNext"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-pNext-pNext"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-src-parameter"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-src-parameter"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-src-parameter  
  `src` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

- <a href="#VUID-VkCopyMicromapToMemoryInfoEXT-mode-parameter"
  id="VUID-VkCopyMicromapToMemoryInfoEXT-mode-parameter"></a>
  VUID-VkCopyMicromapToMemoryInfoEXT-mode-parameter  
  `mode` **must** be a valid
  [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html),
[VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html),
[VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html),
[vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapToMemoryEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMicromapToMemoryInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
