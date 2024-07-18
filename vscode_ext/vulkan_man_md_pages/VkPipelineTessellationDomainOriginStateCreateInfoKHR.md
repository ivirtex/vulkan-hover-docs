# VkPipelineTessellationDomainOriginStateCreateInfo(3) Manual Page

## Name

VkPipelineTessellationDomainOriginStateCreateInfo - Structure specifying
the orientation of the tessellation domain



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineTessellationDomainOriginStateCreateInfo` structure is
defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPipelineTessellationDomainOriginStateCreateInfo {
    VkStructureType               sType;
    const void*                   pNext;
    VkTessellationDomainOrigin    domainOrigin;
} VkPipelineTessellationDomainOriginStateCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance2
typedef VkPipelineTessellationDomainOriginStateCreateInfo VkPipelineTessellationDomainOriginStateCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `domainOrigin` is a
  [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html) value
  controlling the origin of the tessellation domain space.

## <a href="#_description" class="anchor"></a>Description

If the `VkPipelineTessellationDomainOriginStateCreateInfo` structure is
included in the `pNext` chain of
[VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html),
it controls the origin of the tessellation domain. If this structure is
not present, it is as if `domainOrigin` was
`VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT`.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineTessellationDomainOriginStateCreateInfo-sType-sType"
  id="VUID-VkPipelineTessellationDomainOriginStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineTessellationDomainOriginStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO`

- <a
  href="#VUID-VkPipelineTessellationDomainOriginStateCreateInfo-domainOrigin-parameter"
  id="VUID-VkPipelineTessellationDomainOriginStateCreateInfo-domainOrigin-parameter"></a>
  VUID-VkPipelineTessellationDomainOriginStateCreateInfo-domainOrigin-parameter  
  `domainOrigin` **must** be a valid
  [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineTessellationDomainOriginStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
