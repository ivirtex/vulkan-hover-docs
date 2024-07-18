# VK_LUNARG_direct_driver_loading(3) Manual Page

## Name

VK_LUNARG_direct_driver_loading - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

460

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Charles Giessen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_LUNARG_direct_driver_loading%5D%20@charles-lunarg%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_LUNARG_direct_driver_loading%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>charles-lunarg</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_LUNARG_direct_driver_loading](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_LUNARG_direct_driver_loading.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-11-29

**Contributors**  
- Charles Giessen, LunarG

## <a href="#_description" class="anchor"></a>Description

This extension provides a mechanism for applications to add drivers to
the implementation. This allows drivers to be included with an
application without requiring installation and is capable of being used
in any execution environment, such as a process running with elevated
privileges.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingInfoLUNARG.html)

- Extending [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html):

  - [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html)

## <a href="#_new_function_pointers" class="anchor"></a>New Function Pointers

- [PFN_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkGetInstanceProcAddrLUNARG.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDirectDriverLoadingModeLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingModeLUNARG.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkDirectDriverLoadingFlagsLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingFlagsLUNARG.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_LUNARG_DIRECT_DRIVER_LOADING_EXTENSION_NAME`

- `VK_LUNARG_DIRECT_DRIVER_LOADING_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_INFO_LUNARG`

  - `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_LIST_LUNARG`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-11-29 (Charles Giessen)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_LUNARG_direct_driver_loading"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
