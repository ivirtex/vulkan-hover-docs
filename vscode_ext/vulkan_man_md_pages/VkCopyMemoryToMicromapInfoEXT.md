# VkCopyMemoryToMicromapInfoEXT(3) Manual Page

## Name

VkCopyMemoryToMicromapInfoEXT - Parameters for deserializing a micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyMemoryToMicromapInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkCopyMemoryToMicromapInfoEXT {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceOrHostAddressConstKHR    src;
    VkMicromapEXT                    dst;
    VkCopyMicromapModeEXT            mode;
} VkCopyMemoryToMicromapInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `src` is the device or host address to memory containing the source
  data for the copy.

- `dst` is the target micromap for the copy.

- `mode` is a [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html) value
  specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-src-07547"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-src-07547"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-src-07547  
  The source memory pointed to by `src` **must** contain data previously
  serialized using
  [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapToMemoryEXT.html)

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-mode-07548"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-mode-07548"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-mode-07548  
  `mode` **must** be `VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT`

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-src-07549"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-src-07549"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-src-07549  
  The data in `src` **must** have a format compatible with the
  destination physical device as returned by
  [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMicromapCompatibilityEXT.html)

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-dst-07550"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-dst-07550"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-dst-07550  
  `dst` **must** have been created with a `size` greater than or equal
  to that used to serialize the data in `src`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-sType-sType"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-sType-sType"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_MICROMAP_INFO_EXT`

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-pNext-pNext"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-pNext-pNext"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-dst-parameter"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-dst-parameter"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-dst-parameter  
  `dst` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

- <a href="#VUID-VkCopyMemoryToMicromapInfoEXT-mode-parameter"
  id="VUID-VkCopyMemoryToMicromapInfoEXT-mode-parameter"></a>
  VUID-VkCopyMemoryToMicromapInfoEXT-mode-parameter  
  `mode` **must** be a valid
  [VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCopyMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapModeEXT.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToMicromapEXT.html),
[vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToMicromapEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMemoryToMicromapInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
