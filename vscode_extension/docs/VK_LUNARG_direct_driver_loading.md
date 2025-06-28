# VK\_LUNARG\_direct\_driver\_loading(3) Manual Page

## Name

VK\_LUNARG\_direct\_driver\_loading - instance extension



## [](#_registered_extension_number)Registered Extension Number

460

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Charles Giessen [\[GitHub\]charles-lunarg](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_LUNARG_direct_driver_loading%5D%20%40charles-lunarg%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_LUNARG_direct_driver_loading%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_LUNARG\_direct\_driver\_loading](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_LUNARG_direct_driver_loading.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-11-29

**Contributors**

- Charles Giessen, LunarG

## [](#_description)Description

This extension provides a mechanism for applications to add drivers to the implementation. This allows drivers to be included with an application without requiring installation and is capable of being used in any execution environment, such as a process running with elevated privileges.

## [](#_new_structures)New Structures

- [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingInfoLUNARG.html)
- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html):
  
  - [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html)

## [](#_new_function_pointers)New Function Pointers

- [PFN\_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkGetInstanceProcAddrLUNARG.html)

## [](#_new_enums)New Enums

- [VkDirectDriverLoadingModeLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingModeLUNARG.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDirectDriverLoadingFlagsLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingFlagsLUNARG.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_LUNARG_DIRECT_DRIVER_LOADING_EXTENSION_NAME`
- `VK_LUNARG_DIRECT_DRIVER_LOADING_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_INFO_LUNARG`
  - `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_LIST_LUNARG`

## [](#_version_history)Version History

- Revision 1, 2022-11-29 (Charles Giessen)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_LUNARG_direct_driver_loading)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0