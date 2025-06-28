# VkPipelineRepresentativeFragmentTestStateCreateInfoNV(3) Manual Page

## Name

VkPipelineRepresentativeFragmentTestStateCreateInfoNV - Structure specifying representative fragment test



## [](#_c_specification)C Specification

If the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) includes a `VkPipelineRepresentativeFragmentTestStateCreateInfoNV` structure, then that structure includes parameters controlling the representative fragment test.

The `VkPipelineRepresentativeFragmentTestStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_representative_fragment_test
typedef struct VkPipelineRepresentativeFragmentTestStateCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           representativeFragmentTestEnable;
} VkPipelineRepresentativeFragmentTestStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `representativeFragmentTestEnable` controls whether the representative fragment test is enabled.

## [](#_description)Description

If this structure is not included in the `pNext` chain, `representativeFragmentTestEnable` is considered to be `VK_FALSE`, and the representative fragment test is disabled.

If the active fragment shader does not specify the `EarlyFragmentTests` execution mode, the representative fragment shader test has no effect, even if enabled.

Valid Usage (Implicit)

- [](#VUID-VkPipelineRepresentativeFragmentTestStateCreateInfoNV-sType-sType)VUID-VkPipelineRepresentativeFragmentTestStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_representative\_fragment\_test](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_representative_fragment_test.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRepresentativeFragmentTestStateCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0