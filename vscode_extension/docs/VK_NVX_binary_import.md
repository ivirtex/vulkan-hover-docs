# VK\_NVX\_binary\_import(3) Manual Page

## Name

VK\_NVX\_binary\_import - device extension



## [](#_registered_extension_number)Registered Extension Number

30

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_debug\_report

## [](#_contact)Contact

- Eric Werness [\[GitHub\]ewerness-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_binary_import%5D%20%40ewerness-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_binary_import%20extension%2A)
- Liam Middlebrook [\[GitHub\]liam-middlebrook](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_binary_import%5D%20%40liam-middlebrook%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_binary_import%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-11-04

**Contributors**

- Eric Werness, NVIDIA
- Liam Middlebrook, NVIDIA

## [](#_description)Description

This extension allows applications to import CuBIN binaries and execute them.

Note

There is currently no specification language written for this extension. The links to APIs defined by the extension are to stubs that only include generated content such as API declarations and implicit valid usage statements.

## [](#_new_object_types)New Object Types

- [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html)
- [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html)

## [](#_new_commands)New Commands

- [vkCmdCuLaunchKernelNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCuLaunchKernelNVX.html)
- [vkCreateCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCuFunctionNVX.html)
- [vkCreateCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCuModuleNVX.html)
- [vkDestroyCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCuFunctionNVX.html)
- [vkDestroyCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyCuModuleNVX.html)

## [](#_new_structures)New Structures

- [VkCuFunctionCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionCreateInfoNVX.html)
- [VkCuLaunchInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuLaunchInfoNVX.html)
- [VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleCreateInfoNVX.html)
- Extending [VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleCreateInfoNVX.html):
  
  - [VkCuModuleTexturingModeCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleTexturingModeCreateInfoNVX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NVX_BINARY_IMPORT_EXTENSION_NAME`
- `VK_NVX_BINARY_IMPORT_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_CU_FUNCTION_NVX`
  - `VK_OBJECT_TYPE_CU_MODULE_NVX`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_CU_FUNCTION_CREATE_INFO_NVX`
  - `VK_STRUCTURE_TYPE_CU_LAUNCH_INFO_NVX`
  - `VK_STRUCTURE_TYPE_CU_MODULE_CREATE_INFO_NVX`
  - `VK_STRUCTURE_TYPE_CU_MODULE_TEXTURING_MODE_CREATE_INFO_NVX`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_CU_FUNCTION_NVX_EXT`
  - `VK_DEBUG_REPORT_OBJECT_TYPE_CU_MODULE_NVX_EXT`

## [](#_version_history)Version History

- Revision 2, 2024-11-04 (Liam Middlebrook)
  
  - Add [VkCuModuleTexturingModeCreateInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleTexturingModeCreateInfoNVX.html)
- Revision 1, 2021-04-09 (Eric Werness)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NVX_binary_import).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0