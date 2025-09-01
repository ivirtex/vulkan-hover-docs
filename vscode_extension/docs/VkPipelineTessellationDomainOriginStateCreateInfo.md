# VkPipelineTessellationDomainOriginStateCreateInfo(3) Manual Page

## Name

VkPipelineTessellationDomainOriginStateCreateInfo - Structure specifying the orientation of the tessellation domain



## [](#_c_specification)C Specification

The `VkPipelineTessellationDomainOriginStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPipelineTessellationDomainOriginStateCreateInfo {
    VkStructureType               sType;
    const void*                   pNext;
    VkTessellationDomainOrigin    domainOrigin;
} VkPipelineTessellationDomainOriginStateCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance2
typedef VkPipelineTessellationDomainOriginStateCreateInfo VkPipelineTessellationDomainOriginStateCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `domainOrigin` is a [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTessellationDomainOrigin.html) value controlling the origin of the tessellation domain space.

## [](#_description)Description

If the `VkPipelineTessellationDomainOriginStateCreateInfo` structure is included in the `pNext` chain of [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineTessellationStateCreateInfo.html), it controls the origin of the tessellation domain. If this structure is not present, it is as if `domainOrigin` was `VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT`.

Valid Usage (Implicit)

- [](#VUID-VkPipelineTessellationDomainOriginStateCreateInfo-sType-sType)VUID-VkPipelineTessellationDomainOriginStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO`
- [](#VUID-VkPipelineTessellationDomainOriginStateCreateInfo-domainOrigin-parameter)VUID-VkPipelineTessellationDomainOriginStateCreateInfo-domainOrigin-parameter  
  `domainOrigin` **must** be a valid [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTessellationDomainOrigin.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTessellationDomainOrigin.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineTessellationDomainOriginStateCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0